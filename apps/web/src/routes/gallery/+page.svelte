<script lang="ts">
	import { inView } from '$lib/actions/in-view';
	import Figure from '$lib/components/Figure.svelte';
	import { mockPreviews } from '$lib/data/mock-previews';
</script>

<svelte:head>
	<title>Gallery - Open Artwork</title>
	<meta
		name="description"
		content="Open Artwork is an open art data project for exploring artists, artworks, metadata, images, and color palettes. It is primarily a portfolio and learning project for exploring frontend development, backend architecture, and image-processing work."
	/>
</svelte:head>

<section
	class="flex min-h-screen flex-col px-6 py-15 sm:px-10 md:px-15 md:py-20 lg:gap-20 xl:p-25 2xl:p-30"
>
	<div class="group flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto" use:inView>
		<h1
			class="translate-y-5 text-5xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 md:text-6xl"
		>
			Gallery
		</h1>
		<p
			class="translate-y-5 opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
		>
			Lorem Ipsum - Search Art
		</p>
	</div>

	<div class="grid grid-cols-1 gap-15 md:grid-cols-2 lg:grid-cols-3 lg:gap-20 xl:grid-cols-4">
		{#each mockPreviews as preview}
			<a href={`/artwork/${preview.slug}`} use:inView class="group">
				<Figure
					cn="w-full aspect-[3/2] group-hover:scale-102 opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100"
					artwork={{
						title: preview.title,
						artist: preview.primaryArtistName,
						date: preview.dateDisplay
					}}
					img={{ src: preview.imageUrl }}
					showCaption="never"
				/>
			</a>
		{/each}
	</div>
</section>
