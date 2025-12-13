import { z } from 'zod/v4';

export const characterSchema = z.object({
	age: z.string().nullable(),
	anilistId: z.int(),
	birthDay: z.int().nullable(),
	birthMonth: z.int().nullable(),
	birthYear: z.int().nullable(),
	bloodType: z.string().nullable(),
	description: z.string().nullable(),
	gender: z.string().nullable(),
	id: z.string(),
	image: z.string(),
	name: z.string()
});

export type CharacterSchema = z.infer<typeof characterSchema>;
