import { z } from 'zod/v4';

export const loginSchema = z.object({
	username: z.string('Username is required').min(1, 'Username is required'),
	password: z.string('Password is required').min(1, 'Password is required')
});
export const registerSchema = z
	.object({
		email: z.email('Invalid email').min(1, 'Email is required'),
		username: z.string('Username is required').min(1, 'Username is required'),
		password: z
			.string('Password is required')
			.min(8, 'Password must be at least 8 characters long'),
		confirmPassword: z.string('Confirm password is required').min(8, 'Confirm password is required')
	})
	.superRefine((data, ctx) => {
		if (data.confirmPassword != data.password) {
			ctx.addIssue({
				message: 'Passwords do not match',
				path: ['confirmPassword'],
				code: 'custom'
			});
			ctx.addIssue({
				message: 'Passwords do not match',
				path: ['password'],
				code: 'custom'
			});
		}
	});

export type LoginSchema = typeof loginSchema;
export type RegisterSchema = typeof registerSchema;
