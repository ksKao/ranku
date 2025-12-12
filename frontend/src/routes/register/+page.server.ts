import { fail, type Actions } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import { registerSchema } from '../schema';
import { auth } from '$lib/auth';
import { z } from 'zod/v4';

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, zod4(registerSchema));

		if (!form.valid) {
			return fail(400, { form, message: 'Invalid input' });
		}

		try {
			const result = await auth.api.signUpEmail({
				body: {
					email: form.data.email,
					username: form.data.username,
					name: form.data.username,
					password: form.data.password
				},
				asResponse: true
			});

			if (!result.ok) {
				const json = await result.json();

				const err = z
					.object({
						code: z.string(),
						message: z.string()
					})
					.safeParse(json);

				return fail(400, { form, message: err.data?.message ?? 'Failed to register account' });
			}
		} catch (e) {
			if (e instanceof Error) return fail(400, { form, message: e.message });
			else return fail(400, { form, message: 'Failed to register account' });
		}

		return { form };
	}
};
