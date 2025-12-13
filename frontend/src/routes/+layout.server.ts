import { loginSchema } from '$lib/schemas/login.schema';
import { registerSchema } from '$lib/schemas/register.schema';
import { superValidate } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
	return {
		...locals,
		loginForm: await superValidate(zod4(loginSchema)),
		registerForm: await superValidate(zod4(registerSchema))
	};
};
