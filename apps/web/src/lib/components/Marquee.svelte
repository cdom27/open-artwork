<script lang="ts">
	import { inView } from '$lib/actions/in-view';

	export interface MarqueeImage {
		src: string;
		alt?: string;
	}

	interface MarqueeProps {
		images: MarqueeImage[] | string[];
		secondImages?: MarqueeImage[] | string[];
		orientation?: 'horizontal' | 'vertical';
		speed?: number;
		gap?: string;
		itemWidth?: string;
		cn?: string;
	}

	let {
		images,
		secondImages,
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

	const firstLine = $derived(normalizeImages(images));
	const secondLine = $derived(normalizeImages(secondImages ?? images));
	const classes = $derived(
		`opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100 marquee ${orientation} ${cn}`.trim()
	);
	const style = $derived(
		`--marquee-duration: ${Math.max(1, speed)}s; --marquee-gap: ${gap}; --marquee-item-width: ${itemWidth};`
	);
</script>

<div class={classes} {style} use:inView={{ once: false }} aria-label="Scrolling artwork gallery">
	{#if firstLine.length}
		<div class="line">
			<div class="track forward">
				<div class="set">
					{#each firstLine as image}
						<img
							src={image.src}
							alt={image.alt ?? ''}
							loading="lazy"
							class="block h-auto max-w-[80vw] flex-none rounded-sm object-center"
						/>
					{/each}
				</div>
				<div class="set" aria-hidden="true">
					{#each firstLine as image}
						<img
							src={image.src}
							alt=""
							loading="lazy"
							class="block h-auto max-w-[80vw] flex-none rounded-sm object-center"
						/>
					{/each}
				</div>
			</div>
		</div>
	{/if}

	{#if secondLine.length}
		<div class="line">
			<div class="track reverse">
				<div class="set">
					{#each secondLine as image}
						<img
							src={image.src}
							alt={image.alt ?? ''}
							loading="lazy"
							class="block h-auto max-w-[80vw] flex-none rounded-sm object-center"
						/>
					{/each}
				</div>
				<div class="set" aria-hidden="true">
					{#each secondLine as image}
						<img
							src={image.src}
							alt=""
							loading="lazy"
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
		gap: var(--marquee-gap);
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
		overflow: hidden;
	}

	.track {
		display: flex;
		width: max-content;
		animation-duration: var(--marquee-duration);
		animation-timing-function: linear;
		animation-iteration-count: infinite;
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

	.forward {
		animation-name: marquee-forward;
	}

	.reverse {
		animation-name: marquee-reverse;
	}

	@keyframes marquee-forward {
		from {
			transform: translateX(0);
		}

		to {
			transform: translateX(-50%);
		}
	}

	@keyframes marquee-reverse {
		from {
			transform: translateX(-50%);
		}

		to {
			transform: translateX(0);
		}
	}

	.vertical .forward {
		animation-name: marquee-up;
	}

	.vertical .reverse {
		animation-name: marquee-down;
	}

	@keyframes marquee-up {
		from {
			transform: translateY(0);
		}

		to {
			transform: translateY(-50%);
		}
	}

	@keyframes marquee-down {
		from {
			transform: translateY(-50%);
		}

		to {
			transform: translateY(0);
		}
	}

	@media (prefers-reduced-motion: reduce) {
		.track {
			animation: none;
			transform: none;
		}
	}
</style>
