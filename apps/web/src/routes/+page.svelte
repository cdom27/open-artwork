<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import Figure from '$lib/components/Figure.svelte';
	import Link from '$lib/components/Link.svelte';
	import { tools } from '$lib/data/tools';

	const createRings = (
		count: number,
		{ startTop = 80, topStep = 10, startSize = 50, sizeStep = 35 } = {}
	) =>
		Array.from({ length: count }, (_, index) => ({
			top: startTop - index * topStep,
			size: startSize + index * sizeStep
		}));

	const rings = createRings(18);
</script>

<svelte:head>
	<title>Open Artwork</title>
	<meta
		name="description"
		content="Open Artwork is an open art data project for exploring artists, artworks, metadata, images, and color palettes. It is primarily a portfolio and learning project for exploring frontend development, backend architecture, and image-processing work."
	/>
</svelte:head>

<section
	class="grid grid-cols-1 px-2 py-15 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:gap-20 xl:grid-cols-3"
>
	<div class="flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto">
		<h1 class="text-5xl font-normal md:text-6xl">
			Art for<br /> a more open<br /> internet
		</h1>

		<Figure
			cn="lg:hidden"
			imgSrc="/images/samson-mocked-jan-steen-c1670.webp"
			artwork={{ title: 'Samson Mocked', artist: 'Jan Steen', date: 'c. 1670' }}
		/>

		<p>
			Open Artwork is a growing archive of public domain artworks and metadata, with a clean
			developer API and tools for exploring, remixing, and learning from the world&apos;s visual
			culture.
		</p>

		<div class="flex flex-col gap-2.5 pt-5.5 sm:gap-4 xl:flex-row">
			<Link variant="primary" href="/gallery">Explore artworks</Link>
			<Link variant="tertiary" href="/docs">Read API docs</Link>
		</div>
	</div>

	<Figure
		cn="hidden lg:block lg:my-auto xl:col-span-2"
		imgSrc="/images/samson-mocked-jan-steen-c1670.webp"
		artwork={{ title: 'Samson Mocked', artist: 'Jan Steen', date: 'c. 1670' }}
	/>
</section>

<section
	class="relative overflow-clip px-2 py-15 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:gap-20"
>
	<div class="absolute inset-0 flex items-center justify-center">
		{#each rings as ring}
			<div
				class="absolute rounded-full border-2 border-bone-100"
				style={`top: ${ring.top}%; width: calc(${ring.size} * var(--spacing)); height: calc(${ring.size} * var(--spacing));`}
			></div>
		{/each}
	</div>

	<div class="relative z-10 flex flex-col gap-5.5 sm:gap-7.5">
		<h2 class="text-center text-3xl font-normal lg:text-4xl">
			Built from the world's<br /> largest open art collections
		</h2>

		<ul class="flex flex-col gap-5.5 sm:mx-auto sm:flex-row sm:gap-7.5 lg:gap-15">
			<li>
				<div class="flex flex-col items-center gap-1.5">
					<span class="text-5xl">75k+</span><span>artworks</span>
				</div>
			</li>
			<li>
				<div class="flex flex-col items-center gap-1.5">
					<span class="text-5xl">375k+</span><span>color palettes</span>
				</div>
			</li>
		</ul>
	</div>
</section>

<section
	class="relative flex flex-col bg-charcoal-900 text-bone-50 lg:grid lg:grid-cols-2 lg:gap-20"
>
	<div class="flex flex-col gap-5.5 px-2 py-15 sm:px-10 md:px-15 md:py-20">
		<h2 class="text-charcoal-400">The problem</h2>

		<ul class=" flex flex-col gap-5.5 text-4xl font-normal sm:gap-7.5">
			<li>Scattered across institutions</li>
			<li>Inconsistent metadata</li>
			<li>No unified API</li>
			<li>Paywalled access to public domain</li>
			<li>Slow, clunky interfaces</li>
		</ul>
	</div>

	<div
		class="flex flex-col gap-5.5 bg-bone-100 px-2 py-15 text-charcoal-900 sm:gap-7.5 sm:px-10 md:px-15 md:py-20 lg:my-4 lg:w-full lg:rounded-l-sm"
	>
		<div class="aspect-square bg-sky-500 sm:size-80"></div>

		<h2 class="text-charcoal-500">The solution</h2>

		<ul class="flex flex-col gap-5.5 text-4xl font-normal sm:gap-7.5">
			<li>One unified archive</li>
			<li>Color & design resource</li>
			<li>Sharp and intuitive gallery</li>
			<li>Simple developer API</li>
			<li>Clean, consistent metadata</li>
		</ul>
	</div>
</section>

<section class="grid grid-cols-1 bg-bone-100 px-2 py-15 sm:px-10 md:px-15 md:py-20 lg:gap-20">
	<div class="flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto">
		<h2 class="text-4xl font-normal">Other ways to explore</h2>

		<ul
			class="grid grid-cols-1 gap-5.5 pt-5.5 sm:gap-7.5 md:grid-cols-2 lg:grid-cols-3 lg:gap-10 xl:gap-20"
		>
			{#each tools as tool}
				<li>
					<Card
						thumbnail={{ src: tool.thumbnailSrc, alt: '' }}
						title={tool.title}
						description={tool.description}
						link={tool.link}
						variant="primary"
					/>
				</li>
			{/each}
		</ul>
	</div>
</section>
