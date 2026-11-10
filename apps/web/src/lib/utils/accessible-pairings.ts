export type AccessiblePairing = {
	foreground: string;
	background: string;
	ratio: number;
};

const getRelativeLuminance = (hex: string) => {
	const normalized = hex.replace('#', '');
	if (!/^[\da-f]{6}$/i.test(normalized)) return null;

	const channels = [0, 2, 4].map(
		(start) => Number.parseInt(normalized.slice(start, start + 2), 16) / 255
	);
	const linearChannels = channels.map((channel) =>
		channel <= 0.03928 ? channel / 12.92 : Math.pow((channel + 0.055) / 1.055, 2.4)
	);

	return 0.2126 * linearChannels[0] + 0.7152 * linearChannels[1] + 0.0722 * linearChannels[2];
};

export const getAccessiblePairings = (
	colors: string[],
	minimumRatio = 4.5,
	limit = 4
): AccessiblePairing[] => {
	const pairings: AccessiblePairing[] = [];

	for (let firstIndex = 0; firstIndex < colors.length; firstIndex += 1) {
		for (let secondIndex = firstIndex + 1; secondIndex < colors.length; secondIndex += 1) {
			const firstColor = colors[firstIndex];
			const secondColor = colors[secondIndex];
			const firstLuminance = getRelativeLuminance(firstColor);
			const secondLuminance = getRelativeLuminance(secondColor);

			if (firstLuminance === null || secondLuminance === null) continue;

			const ratio =
				(Math.max(firstLuminance, secondLuminance) + 0.05) /
				(Math.min(firstLuminance, secondLuminance) + 0.05);

			if (ratio >= minimumRatio) {
				pairings.push({
					foreground: firstLuminance > secondLuminance ? firstColor : secondColor,
					background: firstLuminance > secondLuminance ? secondColor : firstColor,
					ratio
				});
			}
		}
	}

	return pairings.sort((a, b) => b.ratio - a.ratio).slice(0, limit);
};
