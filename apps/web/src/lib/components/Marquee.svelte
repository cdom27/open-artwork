<script lang="ts">
	import { inView } from '$lib/actions/in-view';
	import { onMount } from 'svelte';

	export interface MarqueeImage {
		src: string;
		alt?: string;
	}

	interface MarqueeProps {
		images: MarqueeImage[] | string[];
		orientation?: 'horizontal' | 'vertical';
		speed?: number;
		gap?: string;
		itemWidth?: string;
		cn?: string;
	}

	let {
		images,
		orientation = 'horizontal',
		speed = 30,
		gap = '1rem',
		itemWidth = '12rem',
		cn = ''
	}: MarqueeProps = $props();

	const normalizeImages = (items: MarqueeImage[] | string[]) =>
		items
			.map((image) => (typeof image === 'string' ? { src: image } : image))
			.filter((image) => image.src);

	let setElement = $state<HTMLDivElement>();
	let distance = $state(0);

	const line = $derived(normalizeImages(images));
	const classes = $derived(
		`opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100 marquee ${orientation} ${cn}`.trim()
	);
	const style = $derived(
		`--marquee-distance: -${distance}px; --marquee-duration: ${Math.max(1, speed * line.length)}s; --marquee-gap: ${gap}; --marquee-item-width: ${itemWidth};`
	);

	const updateDistance = () => {
		if (!setElement) return;

		distance =
			orientation === 'vertical'
				? setElement.getBoundingClientRect().height
				: setElement.getBoundingClientRect().width;
	};

	onMount(() => {
		updateDistance();

		if (!setElement) return;
		const resizeObserver = new ResizeObserver(updateDistance);
		resizeObserver.observe(setElement);

		return () => resizeObserver.disconnect();
	});
</script>

<div class={classes} {style} use:inView={{ once: false }} aria-label="Scrolling artwork gallery">
	{#if line.length}
		<div class="line">
			<div class="track">
				<div class="set" bind:this={setElement}>
					{#each line as image}
						<img
							src={image.src}
							alt={image.alt ?? ''}
							loading="eager"
							class="block h-auto max-w-[80vw] flex-none rounded-sm object-center"
						/>
					{/each}
				</div>
				<div class="set" aria-hidden="true">
					{#each line as image}
						<img
							src={image.src}
							alt=""
							loading="eager"
							class="block h-auto max-w-[80vw] flex-none rounded-sm object-center"
						/>
					{/each}
				</div>
			</div>
		</div>
	{/if}
</div>

<style>
	img {
		width: var(--marquee-item-width);
	}

	.marquee {
		display: flex;
		overflow: hidden;
		contain: content;
	}

	.marquee.horizontal {
		flex-direction: column;
	}

	.marquee.vertical {
		flex-direction: row;
	}

	.line {
		position: relative;
		overflow: hidden;
		flex: 1 1 0%;
		min-width: 0;
		min-height: 0;
	}

	.horizontal .line {
		width: 100%;
	}

	.vertical .line {
		height: 100%;
		display: flex;
		justify-content: center;
	}

	.track {
		display: flex;
		width: max-content;
		animation: marquee-horizontal var(--marquee-duration) linear infinite;
		animation-play-state: paused;
		will-change: transform;
	}

	.in-view .track {
		animation-play-state: running;
	}

	.set {
		display: flex;
		flex: none;
		gap: var(--marquee-gap);
	}

	.horizontal .set {
		padding-right: var(--marquee-gap);
	}

	.vertical .track {
		flex-direction: column;
		width: 100%;
		align-items: center;
		animation-name: marquee-vertical;
	}

	.vertical .set {
		flex-direction: column;
		padding-bottom: var(--marquee-gap);
		align-items: center;
		width: 100%;
	}

	.vertical img {
		max-width: 100%;
	}

	@keyframes marquee-horizontal {
		from {
			transform: translateX(0);
		}

		to {
			transform: translateX(var(--marquee-distance));
		}
	}

	@keyframes marquee-vertical {
		from {
			transform: translateY(0);
		}

		to {
			transform: translateY(var(--marquee-distance));
		}
	}

	@media (prefers-reduced-motion: reduce) {
		.track {
			animation: none;
			transform: none;
		}
	}
</style>
