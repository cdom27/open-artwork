<script>
	import Figure from '$lib/components/Figure.svelte';
	import Link from '$lib/components/Link.svelte';
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
