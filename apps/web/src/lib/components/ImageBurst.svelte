<script lang="ts">
	import { onMount } from 'svelte';

	export interface BurstImage {
		src: string;
		alt?: string;
	}

	interface Burst {
		id: number;
		image: BurstImage;
		left: number;
		top: number;
		zIndex: number;
		size: number;
		driftX: number;
		driftY: number;
		delay: number;
	}

	type ClusterTier = 'small' | 'medium' | 'large';

	interface ClusterSlot {
		left: number;
		top: number;
		spread: number;
		tier: ClusterTier;
	}

	interface ImageBurstProps {
		images: BurstImage[] | string[];
		spawnInterval?: number;
		lifetime?: number;
		minTargetSize?: number;
		maxTargetSize?: number;
		minVerticalDrift?: number;
		maxVerticalDrift?: number;
		minHorizontalDrift?: number;
		maxHorizontalDrift?: number;
		maxVisible?: number;
		backdropSize?: number;
		backdropDriftX?: number;
		backdropDriftY?: number;
		backdropMotionDuration?: number;
		compositionOffsetX?: number;
		compositionOffsetY?: number;
	}

	const clusterSlots: ClusterSlot[] = [
		{ left: 50, top: 18, spread: 4, tier: 'medium' },
		{ left: 75, top: 24, spread: 5, tier: 'small' },
		{ left: 84, top: 46, spread: 5, tier: 'medium' },
		{ left: 76, top: 70, spread: 5, tier: 'large' },
		{ left: 57, top: 83, spread: 4, tier: 'small' },
		{ left: 31, top: 78, spread: 5, tier: 'medium' },
		{ left: 16, top: 59, spread: 5, tier: 'small' },
		{ left: 22, top: 34, spread: 4, tier: 'large' },
		{ left: 40, top: 29, spread: 5, tier: 'medium' },
		{ left: 64, top: 57, spread: 5, tier: 'large' }
	];

	let {
		images,
		spawnInterval = 200,
		lifetime = 2000,
		minTargetSize = 40,
		maxTargetSize = 68,
		minVerticalDrift = 3,
		maxVerticalDrift = 8,
		minHorizontalDrift = 30,
		maxHorizontalDrift = 40,
		maxVisible = 10,
		backdropSize = 100,
		backdropDriftX = 3,
		backdropDriftY = 2,
		backdropMotionDuration = 9000,
		compositionOffsetX = -4,
		compositionOffsetY = 0
	}: ImageBurstProps = $props();

	let bursts = $state<Burst[]>([]);
	let nextId = 0;
	let imageIndex = 0;
	let slotIndex = 0;
	let spawnTimer: ReturnType<typeof setInterval> | undefined;
	const removalTimers = new Map<number, ReturnType<typeof setTimeout>>();

	const randomBetween = (min: number, max: number) => min + Math.random() * (max - min);

	const normalizedImages = $derived(
		images
			.map((image) => (typeof image === 'string' ? { src: image } : image))
			.filter((image) => image.src)
	);

	const getTargetSize = (tier: ClusterTier) => {
		const multipliers = {
			small: [0.3, 0.45],
			medium: [0.45, 0.65],
			large: [0.65, 0.9]
		};
		const [minMultiplier, maxMultiplier] = multipliers[tier];

		return randomBetween(minTargetSize * minMultiplier, maxTargetSize * maxMultiplier);
	};

	const spawn = () => {
		if (!normalizedImages.length) return;

		const image = normalizedImages[imageIndex % normalizedImages.length];
		const slot = clusterSlots[slotIndex % clusterSlots.length];
		imageIndex += 1;
		slotIndex += 1;

		const id = nextId++;
		const delay = randomBetween(0, 120);
		if (bursts.length >= maxVisible) {
			const oldestBurst = bursts[0];
			const oldestTimer = removalTimers.get(oldestBurst.id);

			if (oldestTimer) clearTimeout(oldestTimer);
			removalTimers.delete(oldestBurst.id);
			bursts = bursts.slice(1);
		}

		bursts = [
			...bursts,
			{
				id,
				image,
				left: slot.left + randomBetween(-slot.spread, slot.spread),
				top: slot.top + randomBetween(-slot.spread, slot.spread),
				zIndex: slot.tier === 'large' ? 3 : slot.tier === 'medium' ? 2 : 1,
				size: getTargetSize(slot.tier),
				driftX: randomBetween(minHorizontalDrift, maxHorizontalDrift),
				driftY: randomBetween(minVerticalDrift, maxVerticalDrift),
				delay
			}
		];

		removalTimers.set(
			id,
			setTimeout(
				() => {
					bursts = bursts.filter((burst) => burst.id !== id);
					removalTimers.delete(id);
				},
				delay + lifetime + 50
			)
		);
	};

	onMount(() => {
		if (!normalizedImages.length) return;

		spawn();
		spawnTimer = setInterval(spawn, Math.max(16, spawnInterval));

		return () => {
			if (spawnTimer) clearInterval(spawnTimer);
			for (const timer of removalTimers.values()) clearTimeout(timer);
			removalTimers.clear();
		};
	});
</script>

<div class="relative isolate aspect-square w-full" aria-hidden="true">
	{#if normalizedImages.length}
		<div
			class="backdrop"
			style={`--backdrop-size: ${backdropSize}%; --backdrop-drift-x: ${backdropDriftX}%; --backdrop-drift-y: ${backdropDriftY}%; --backdrop-motion-duration: ${backdropMotionDuration}ms;`}
		>
			<img
				src={normalizedImages[0].src}
				alt=""
				class="block h-auto w-full rounded-sm saturate-[0.85]"
			/>
		</div>
	{/if}

	{#each bursts as burst (burst.id)}
		<div
			class="burst"
			style={`--left: ${burst.left}%; --top: ${burst.top}%; --offset-x: ${compositionOffsetX}%; --offset-y: ${compositionOffsetY}%; --z-index: ${burst.zIndex}; --size: ${burst.size}%; --drift-x: ${burst.driftX}px; --drift-y: ${burst.driftY}px; --delay: ${burst.delay}ms; --lifetime: ${lifetime}ms;`}
		>
			<img
				src={burst.image.src}
				alt={burst.image.alt ?? ''}
				class="block h-auto w-full rounded-sm saturate-[0.85]"
			/>
		</div>
	{/each}
</div>

<style>
	.backdrop {
		position: absolute;
		top: 50%;
		left: 50%;
		z-index: 0;
		width: var(--backdrop-size);
		height: auto;
		transform: translate(-50%, -50%);
		animation: backdrop-float var(--backdrop-motion-duration) ease-in-out infinite alternate;
		animation-play-state: var(--burst-animation-state, paused);
		will-change: transform;
	}

	.burst {
		position: absolute;
		top: calc(var(--top) + var(--offset-y));
		left: calc(var(--left) + var(--offset-x));
		z-index: var(--z-index);
		width: var(--size);
		height: auto;
		transform: translate(-50%, -50%) scale(1);
		animation: burst-shrink var(--lifetime) cubic-bezier(0.22, 1, 0.36, 1) var(--delay) forwards;
		animation-play-state: var(--burst-animation-state, paused);
		will-change: transform, opacity;
	}

	@keyframes backdrop-float {
		0% {
			transform: translate(
					calc(-50% - var(--backdrop-drift-x)),
					calc(-50% + var(--backdrop-drift-y))
				)
				scale(1);
		}

		100% {
			transform: translate(
					calc(-50% + var(--backdrop-drift-x)),
					calc(-50% - var(--backdrop-drift-y))
				)
				scale(1);
		}
	}

	@keyframes burst-shrink {
		0% {
			transform: translate(-50%, -50%) scale(1);
		}

		50% {
			transform: translate(calc(-50% + var(--drift-x)), calc(-50% - var(--drift-y))) scale(1);
		}

		92% {
			opacity: 1;
			transform: translate(calc(-50% + var(--drift-x)), calc(-50% - var(--drift-y))) scale(0.35);
		}

		100% {
			opacity: 0;
			transform: translate(calc(-50% + var(--drift-x)), calc(-50% - var(--drift-y))) scale(0.01);
		}
	}

	@media (prefers-reduced-motion: reduce) {
		.backdrop {
			animation: none;
		}

		.burst {
			animation: none;
			opacity: 1;
			transform: translate(-50%, -50%) scale(1);
		}
	}
</style>
