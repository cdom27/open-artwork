export type Preview = {
	title: string;
	slug: string;
	dateDisplay: string;
	primaryArtistName: string;
	imageUrl: string;
};

export type PreviewPage = {
	artworks: Preview[];
	totalResults: number;
	pageSize: number;
	page: number;
	hasMore: boolean;
};
