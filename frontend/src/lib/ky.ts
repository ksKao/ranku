import { env } from '$env/dynamic/public';
import ky from 'ky';
import { authClient } from './auth-client';

export const kyClient = ky.create({
	prefixUrl: env.PUBLIC_BACKEND_URL
});

export const kyClientAuthed = kyClient.extend({
	hooks: {
		beforeRequest: [
			async (request, _, { retryCount }) => {
				if (retryCount === 0) {
					const { data } = await authClient.token();

					if (data) {
						request.headers.set('Authorization', `Bearer ${data.token}`);
					}
				}
			}
		]
	}
});
