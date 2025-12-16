import { auth } from '$lib/auth';
import { svelteKitHandler } from 'better-auth/svelte-kit';
import { building } from '$app/environment';

export async function handle({ event, resolve }) {
	// Fetch current session from Better Auth
	const session = await auth.api.getSession({
		headers: event.request.headers
	});

	// Make session and user available on server
	if (session) {
		event.locals.session = session.session;
		event.locals.user = session.user;

		try {
			const data = await auth.api.getToken({ headers: event.request.headers });

			event.locals.token = data.token;
		} catch (e) {
			console.log('Failed to get token: ', e);
		}
	}

	return svelteKitHandler({ event, resolve, auth, building });
}
