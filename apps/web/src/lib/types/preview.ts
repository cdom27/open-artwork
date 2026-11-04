export type Preview = {
	title: string;
	slug: string;
	dateDisplay: string;
	primaryArtistName: string;
	thumbnailUrl: string;
};

export type PreviewPage = {
	artworks: Preview[];
	totalResults: number;
	pageSize: number;
	page: number;
	hasMore: boolean;
};
