import { z } from 'zod/v4';

export const leaderboardSchema = z.array(
	z.object({
		id: z.string(),
		name: z.string(),
		image: z.string(),
		for: z.int(),
		against: z.int(),
		score: z.int()
	})
);

export type LeaderboardSchema = z.infer<typeof leaderboardSchema>;
