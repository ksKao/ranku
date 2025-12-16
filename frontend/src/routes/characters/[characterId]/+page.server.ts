import { kyClient } from '$lib/ky';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { z } from 'zod/v4';
import { characterSchema } from '$lib/schemas/character.schema';

export const load: PageServerLoad = async ({ params }) => {
	try {
		const response = await kyClient.get(`characters/${params.characterId}`).json();

		const character = z
			.object({
				...characterSchema.shape,
				animes: z.array(z.string()),
				likes: z.number(),
				liked: z.boolean()
			})
			.omit({
				anime: true
			})
			.parse(response);

		return {
			a: 'b',
			character
		};
	} catch {
		error(404);
	}
};
