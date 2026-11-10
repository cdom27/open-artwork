export const artworkMoods = ['Colorful', 'Bright', 'Muted', 'Deep', 'Dark', 'Original'] as const;
export type ArtworkMood = (typeof artworkMoods)[number];

export type ArtworkDetails = {
	title: string;
	objectType: string;
	department: string;
	classification: string;
	culture: string;
	period: string;
	dateDisplay: string;
	medium: string;
	dimensions: string;
	tags: string[];
	constituents: string[];
	creditLine: string;
	palettes: ArtworkPalette[];
	fullUrl: string;
};

export type ArtworkPalette = {
	mood: ArtworkMood;
	composition: string;
	colors: {
		hex: string;
	}[];
};
