import { z } from 'zod/v4';
import { characterSchema } from './character.schema';

export const matchupSchema = z.object({
	char1: characterSchema,
	char2: characterSchema
});

export type MatchupSchema = z.infer<typeof matchupSchema>;
