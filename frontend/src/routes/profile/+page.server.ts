import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	if (!locals.session || !locals.user) {
		redirect(301, '/');
	}

	return { ...locals };
};
