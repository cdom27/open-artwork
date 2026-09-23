import type { Action } from 'svelte/action';

interface InViewOptions {
	threshold?: number;
	once?: boolean;
}

interface InViewAttributes {
	'on:enterviewport'?: (e: CustomEvent<void>) => void;
}

export const inView: Action<HTMLElement, InViewOptions | undefined, InViewAttributes> = (
	node,
	options = {}
) => {
	const { threshold = 0.2, once = true } = options;

	const observer = new IntersectionObserver(
		([entry]) => {
			if (entry.isIntersecting) {
				node.classList.add('in-view');
				node.dispatchEvent(new CustomEvent('enterviewport'));
				if (once) observer.unobserve(node);
			} else if (!once) {
				node.classList.remove('in-view');
			}
		},
		{ threshold }
	);

	observer.observe(node);

	return {
		destroy() {
			observer.disconnect();
		}
	};
};
