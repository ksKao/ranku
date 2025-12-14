import type { PageServerLoad } from './$types';
import { kyClient } from '$lib/ky';
import { leaderboardSchema } from '$lib/schemas/leaderboard.schema';

export const load: PageServerLoad = async () => {
	const json = await kyClient.get('leaderboard').json();

	const parsed = leaderboardSchema.parse(json);

	return { leaderboard: parsed };
};
