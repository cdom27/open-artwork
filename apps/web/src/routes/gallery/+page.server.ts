import { env } from '$env/dynamic/private';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { ApiResponse } from '$lib/types/api';
import type { PreviewPage } from '$lib/types/preview';

const apiBaseUrl = env.API_BASE_URL ?? 'http://localhost:8080';

export const load: PageServerLoad = async ({ fetch, url }) => {
	const apiUrl = new URL('/v0/api/artworks', apiBaseUrl);
	apiUrl.search = url.search;

	const response = await fetch(apiUrl);

	if (!response.ok) {
		error(response.status, 'Unable to load artworks.');
	}

	const body = (await response.json()) as ApiResponse<PreviewPage>;

	return body.data;
};
