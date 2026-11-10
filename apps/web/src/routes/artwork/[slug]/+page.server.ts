import { env } from '$env/dynamic/private';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { ApiResponse } from '$lib/types/api';
import type { ArtworkDetails } from '$lib/types/artwork';

const apiBaseUrl = env.API_BASE_URL ?? 'http://localhost:8080';

export const load: PageServerLoad = async ({ fetch, params }) => {
	const slug = encodeURIComponent(params.slug);
	const [response, relatedResponse] = await Promise.all([
		fetch(new URL(`/v0/api/artworks/${slug}`, apiBaseUrl)),
		fetch(new URL(`/v0/api/artworks/${slug}/related`, apiBaseUrl))
	]);

	if (!response.ok) {
		error(response.status, 'Unable to load artwork.');
	}
	if (!relatedResponse.ok) {
		error(relatedResponse.status, 'Unable to load related artworks.');
	}

	const body = (await response.json()) as ApiResponse<ArtworkDetails>;
	const relatedBody = (await relatedResponse.json()) as ApiResponse<{
		artworks: ArtworkDetails['relatedArtworks'];
	}>;

	return { ...body.data, relatedArtworks: relatedBody.data.artworks };
};
