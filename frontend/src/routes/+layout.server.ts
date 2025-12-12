import { superValidate } from 'sveltekit-superforms';
import type { LayoutServerLoad } from './$types';
import { zod4 } from 'sveltekit-superforms/adapters';
import { loginSchema } from './schema';

export const load: LayoutServerLoad = async ({ locals }) => {
	return { ...locals, loginForm: await superValidate(zod4(loginSchema)) };
};
