import { superValidate } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import type { LayoutServerLoad } from './$types';
import { loginSchema, registerSchema } from './schema';

export const load: LayoutServerLoad = async ({ locals }) => {
	return {
		...locals,
		loginForm: await superValidate(zod4(loginSchema)),
		registerForm: await superValidate(zod4(registerSchema))
	};
};
