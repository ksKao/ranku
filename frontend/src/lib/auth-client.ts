import { jwtClient, usernameClient } from 'better-auth/client/plugins';
import { createAuthClient } from 'better-auth/svelte';

export const authClient = createAuthClient({
	plugins: [usernameClient(), jwtClient()]
});

export type Session = typeof authClient.$Infer.Session;
