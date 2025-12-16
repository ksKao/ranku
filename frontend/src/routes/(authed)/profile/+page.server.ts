import { kyClient } from '$lib/ky';
import { characterSchema } from '$lib/schemas/character.schema';
import { z } from 'zod/v4';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	try {
		const response = await kyClient
			.get('likes', {
				headers: {
					Authorization: `Bearer ${locals.token}`
				}
			})
			.json();

		const likedCharacters = z
			.array(
				characterSchema.omit({
					anime: true
				})
			)
			.parse(response);

		return {
			likedCharacters
		};
	} catch (e) {
		console.error('Failed to fetch liked characters: ', e);
		return {
			likedCharacters: null
		};
	}
};
