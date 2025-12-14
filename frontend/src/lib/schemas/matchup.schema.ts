import { z } from 'zod/v4';
import { characterSchema } from './character.schema';

export const matchupSchema = z.object({
	char1: characterSchema.nullable(),
	char2: characterSchema.nullable()
});

export type MatchupSchema = z.infer<typeof matchupSchema>;
