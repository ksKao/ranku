import { fail, type Actions } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { loginSchema } from '../schema';
import { auth } from '$lib/auth';

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, zod4(loginSchema));

		if (!form.valid) {
			return fail(400, { form, message: 'Invalid input' });
		}

		try {
			const result = await auth.api.signInUsername({
				body: {
					...form.data
				}
			});

			if (!result) {
				return fail(400, { message: 'Invalid credentials' });
			}
		} catch {
			return fail(400, { form, message: 'Invalid credentials' });
		}

		return { form };
	}
};
