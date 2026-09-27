<script lang="ts">
	interface RingsProps {
		amount?: number;
		spacing?: number;
	}

	let { amount = 18, spacing = 52 }: RingsProps = $props();

	const rings = $derived(
		Array.from({ length: Math.max(0, Math.floor(amount)) }, (_, index) => {
			const radius = 90 + index * spacing;
			const centerX = 500;
			const centerY = 560 + index * 8;
			const startX = Math.max(0, centerX - radius);
			const endX = Math.min(1000, centerX + radius);
			const startY = centerY - Math.sqrt(radius ** 2 - (startX - centerX) ** 2);
			const endY = centerY - Math.sqrt(radius ** 2 - (endX - centerX) ** 2);

			return {
				d: `M ${startX} ${startY} A ${radius} ${radius} 0 0 1 ${endX} ${endY}`,
				delay: index * 5
			};
		})
	);
</script>

<svg
	class="rings pointer-events-none absolute inset-0 h-full w-full"
	viewBox="0 0 1000 500"
	preserveAspectRatio="xMidYMid slice"
	aria-hidden="true"
>
	{#each rings as ring}
		<path
			class="ring"
			d={ring.d}
			style={`--ring-delay: ${ring.delay}ms`}
			fill="none"
			stroke="currentColor"
			stroke-width="1.5"
		/>
	{/each}
</svg>

<style>
	.rings {
		color: var(--color-bone-100);
		contain: paint;
	}

	.ring {
		opacity: 0;
		transition: opacity 600ms;
		transition-delay: var(--ring-delay);
	}

	:global(.in-view) .ring {
		opacity: 1;
	}

	@media (prefers-reduced-motion: reduce) {
		.ring {
			transition: none;
			opacity: 1;
		}
	}
</style>
