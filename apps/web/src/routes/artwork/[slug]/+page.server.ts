import { env } from '$env/dynamic/private';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { ApiResponse } from '$lib/types/api';
import type { ArtworkDetails } from '$lib/types/artwork';

const apiBaseUrl = env.API_BASE_URL ?? 'http://localhost:8080';

export const load: PageServerLoad = async ({ fetch, params }) => {
	const apiUrl = new URL(`/v0/api/artworks/${encodeURIComponent(params.slug)}`, apiBaseUrl);
	const response = await fetch(apiUrl);

	if (!response.ok) {
		error(response.status, 'Unable to load artwork.');
	}

	const body = (await response.json()) as ApiResponse<ArtworkDetails>;

	return body.data;
};
