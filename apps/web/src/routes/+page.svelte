<script lang="ts">
	import { inView } from '$lib/actions/in-view';
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

	let artworksCount = $state(0);
	let palettesCount = $state(0);

	const animateCount = (target: number, update: (value: number) => void, duration = 1000) => {
		const start = performance.now();

		const tick = (now: number) => {
			const progress = Math.min((now - start) / duration, 1);
			const easedProgress = 1 - Math.pow(1 - progress, 3);
			update(Math.round(target * easedProgress));

			if (progress < 1) requestAnimationFrame(tick);
		};

		requestAnimationFrame(tick);
	};

	const startCounters = () => {
		animateCount(75, (value) => (artworksCount = value));
		animateCount(375, (value) => (palettesCount = value));
	};
</script>

<svelte:head>
	<title>Open Artwork</title>
	<meta
		name="description"
		content="Open Artwork is an open art data project for exploring artists, artworks, metadata, images, and color palettes. It is primarily a portfolio and learning project for exploring frontend development, backend architecture, and image-processing work."
	/>
</svelte:head>

<section
	class="group grid grid-cols-1 px-2 py-15 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:gap-20 xl:grid-cols-3 xl:p-25 2xl:p-30"
	use:inView
>
	<div class="flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto">
		<h1
			class="translate-y-5 text-5xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 md:text-6xl"
		>
			Art for<br /> a more open<br /> internet
		</h1>

		<Figure
			cn="lg:hidden opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100"
			img={{
				src: '/images/samson-mocked-jan-steen-c1670.webp',
				width: 1023,
				height: 724,
				loading: 'eager'
			}}
			artwork={{ title: 'Samson Mocked', artist: 'Jan Steen', date: 'c. 1670' }}
		/>

		<p
			class="translate-y-5 opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
		>
			Open Artwork is a growing archive of public domain artworks and metadata, with a clean
			developer API and tools for exploring, remixing, and learning from the world&apos;s visual
			culture.
		</p>

		<div
			class="flex translate-y-5 flex-col gap-2.5 pt-5.5 opacity-0 transition-all delay-240 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 sm:gap-4 xl:flex-row"
		>
			<Link variant="primary" href="/gallery">Explore artworks</Link>
			<Link variant="tertiary" href="/docs">Read API docs</Link>
		</div>
	</div>

	<Figure
		cn="hidden lg:block lg:my-auto xl:col-span-2 opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100"
		img={{
			src: '/images/samson-mocked-jan-steen-c1670.webp',
			width: 1023,
			height: 724,
			loading: 'eager'
		}}
		artwork={{ title: 'Samson Mocked', artist: 'Jan Steen', date: 'c. 1670' }}
	/>
</section>

<section
	class="group relative overflow-clip px-2 py-15 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:gap-20 xl:p-25 2xl:p-30"
	use:inView
	onenterviewport={startCounters}
>
	<div class="absolute inset-0 flex items-center justify-center">
		{#each rings as ring, i}
			<div
				class="absolute rounded-full border-2 border-bone-100 opacity-0 transition-opacity delay-[calc(var(--i)*5ms)] duration-600 group-[&.in-view]:opacity-100"
				style={`--i: ${i}; top: ${ring.top}%; width: calc(${ring.size} * var(--spacing)); height: calc(${ring.size} * var(--spacing));`}
			></div>
		{/each}
	</div>

	<div class="relative z-10 flex flex-col gap-5.5 sm:gap-7.5">
		<h2
			class="translate-y-5 text-center text-3xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 lg:text-4xl"
		>
			Built from the world's<br /> largest open art collections
		</h2>

		<ul class="flex flex-col gap-5.5 sm:mx-auto sm:flex-row sm:gap-7.5 lg:gap-15">
			<li>
				<div
					class="flex translate-y-5 flex-col items-center gap-1.5 opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
				>
					<span class="text-5xl">{artworksCount}k+</span><span>artworks</span>
				</div>
			</li>
			<li>
				<div
					class="flex translate-y-5 flex-col items-center gap-1.5 opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
				>
					<span class="text-5xl">{palettesCount}k+</span><span>color palettes</span>
				</div>
			</li>
		</ul>
	</div>
</section>

<section
	class="group relative flex flex-col overflow-clip bg-charcoal-900 text-bone-50 lg:grid lg:grid-cols-2 lg:gap-20"
	use:inView
>
	<div
		class="group flex flex-col gap-5.5 px-2 py-15 sm:px-10 md:px-15 md:py-20 xl:p-25 2xl:p-30"
		use:inView
	>
		<h2
			class="translate-y-5 text-charcoal-400 opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
		>
			The problem
		</h2>

		<ul class="flex flex-col gap-5.5 text-4xl font-normal sm:gap-7.5">
			<li
				class="translate-y-5 opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Scattered across institutions
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-240 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Inconsistent metadata
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-320 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				No unified API
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-400 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Paywalled access to public domain
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-480 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Slow, clunky interfaces
			</li>
		</ul>
	</div>

	<div
		class="group flex flex-col gap-5.5 bg-bone-100 px-2 py-15 text-charcoal-900 sm:gap-7.5 sm:px-10 md:px-15 md:py-20 lg:my-4 lg:w-full lg:translate-x-80 lg:rounded-l-sm lg:transition-all lg:delay-80 lg:duration-400 lg:group-[&.in-view]:translate-x-0 xl:p-25 2xl:p-30"
		use:inView
	>
		<div
			class="aspect-square bg-sky-500 opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100 sm:size-80"
		></div>

		<h2
			class="translate-y-5 text-charcoal-500 opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
		>
			The solution
		</h2>

		<ul class="flex flex-col gap-5.5 text-4xl font-normal sm:gap-7.5">
			<li
				class="translate-y-5 opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				One unified archive
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-240 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Color & design resource
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-320 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Sharp and intuitive gallery
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-400 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Simple developer API
			</li>
			<li
				class="translate-y-5 opacity-0 transition-all delay-480 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Clean, consistent metadata
			</li>
		</ul>
	</div>
</section>

<section
	class="group grid grid-cols-1 bg-bone-100 px-2 py-15 sm:px-10 md:px-15 md:py-20 xl:p-25 2xl:p-30"
	use:inView
>
	<div class="flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto">
		<h2
			class="translate-y-5 text-4xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
		>
			Other ways to explore
		</h2>

		<ul
			class="group grid grid-cols-1 gap-5.5 pt-5.5 sm:gap-7.5 md:grid-cols-2 lg:grid-cols-3 lg:gap-10"
			use:inView
		>
			{#each tools as tool}
				<li>
					<Card
						thumbnail={{
							src: tool.thumbnailSrc,
							width: 1023,
							height: 724,
							alt: 'Placeholder image',
							loading: 'lazy'
						}}
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

<section
	class="group grid grid-cols-1 bg-charcoal-900 px-2 py-15 text-stone-50 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:gap-20 xl:gap-30 xl:p-25 2xl:gap-100 2xl:p-30"
	use:inView
>
	<div class="flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto">
		<h2
			class="translate-y-5 text-4xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
		>
			Download the data
		</h2>

		<Figure
			cn="lg:hidden sm:h-150 sm:mx-auto"
			img={{
				src: '/images/bust-of-a-youth-francesco-mochi-1640.webp',
				width: 673,
				height: 841,
				loading: 'lazy'
			}}
			artwork={{
				title: 'Bust of a Youth (perhaps Saint John the Baptist)',
				artist: 'Francesco Mochi',
				date: '1630-1640'
			}}
		/>

		<p
			class="translate-y-5 opacity-0 transition-all delay-240 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
		>
			Our dataset of public domain artwork metadata is available for anyone to use. Download the
			full dataset as CSV from our GitHub repository. Images and media files themselves are not
			included due to their size.
		</p>

		<div class="flex flex-col gap-2.5 pt-5.5 sm:flex-row sm:gap-4">
			<Link
				variant="secondary"
				href="/gallery"
				cn="translate-y-5 opacity-0 transition-all delay-320 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
				>Explore artworks</Link
			>
			<Link
				variant="tertiary"
				href="https://github.com/cdom27/open-artwork"
				target="_blank"
				rel="noopener noreferrer"
				cn="translate-y-5 opacity-0 transition-all delay-320 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
				>View the data</Link
			>
		</div>

		<div class="flex flex-col gap-5.5 pt-10 sm:gap-5.5">
			<h3
				class="translate-y-5 text-3xl font-normal opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Notes for users
			</h3>
			<p
				class="translate-y-5 opacity-0 transition-all delay-240 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				The collection is built from public-domain artworks and metadata made available through
				open-access collections, APIs, and datasets. Open Artwork does not claim ownership of the
				underlying artworks or source data.
			</p>
			<p
				class="translate-y-5 opacity-0 transition-all delay-320 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				Rights information is provided by the source institution. We encourage users to credit the
				appropriate institution and source for each image.
			</p>
		</div>
	</div>

	<Figure
		cn="hidden lg:block h-full w-full"
		img={{
			src: '/images/bust-of-a-youth-francesco-mochi-1640.webp',
			width: 673,
			height: 841,
			loading: 'lazy'
		}}
		artwork={{
			title: 'Bust of a Youth (perhaps Saint John the Baptist)',
			artist: 'Francesco Mochi',
			date: '1630-1640'
		}}
	/>
</section>
