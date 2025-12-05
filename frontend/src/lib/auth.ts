import { getRequestEvent } from '$app/server';
import { DB_CONNECTION_STRING } from '$env/static/private';
import { betterAuth } from 'better-auth';
import { jwt, username } from 'better-auth/plugins';
import { sveltekitCookies } from 'better-auth/svelte-kit';
import { Pool } from 'pg';

export const auth = betterAuth({
	database: new Pool({
		connectionString: DB_CONNECTION_STRING
	}),
	emailAndPassword: {
		enabled: true
	},
	plugins: [jwt(), username(), sveltekitCookies(getRequestEvent)] // make sure this is the last plugin in the array
});
