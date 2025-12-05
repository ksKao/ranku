import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals, fetch }) => {
	if (!locals.session?.token) redirect(302, '/login');

	const response = await fetch(`/api/auth/token`, {
		headers: {
			Authorization: `Bearer ${locals.session.token}`
		}
	});

	if (!response.ok) redirect(302, '/login');

	const json: unknown = await response.json();

	if (json && typeof json === 'object' && 'token' in json && typeof json.token === 'string') {
		return { ...locals, token: json.token };
	}

	redirect(302, '/login');
};
