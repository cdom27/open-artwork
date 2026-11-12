export const getDisplayHex = (hex: string) => hex.replace('#', '').toUpperCase();

export const handleCopyColor = async (hex: string, onCopied?: () => void) => {
	try {
		await navigator.clipboard.writeText(hex);
		onCopied?.();
	} catch (error) {
		console.error('Unable to copy color to clipboard.', error);
	}
};

export const getTextColor = (hex: string) => {
	const normalized = hex.replace('#', '');
	if (normalized.length !== 6) return 'text-bone-50';

	const red = Number.parseInt(normalized.slice(0, 2), 16);
	const green = Number.parseInt(normalized.slice(2, 4), 16);
	const blue = Number.parseInt(normalized.slice(4, 6), 16);
	const luminance = (0.299 * red + 0.587 * green + 0.114 * blue) / 255;

	return luminance > 0.6 ? 'text-charcoal-900' : 'text-bone-50';
};
