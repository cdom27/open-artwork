<script lang="ts">
	import { page } from '$app/state';
	import { inView } from '$lib/actions/in-view';
	import Button from '$lib/components/Button.svelte';
	import Figure from '$lib/components/Figure.svelte';
	import Link from '$lib/components/Link.svelte';
	import type { PreviewPage } from '$lib/types/preview';

	let { data }: { data: PreviewPage } = $props();

	const totalPages = $derived(Math.ceil(data.totalResults / data.pageSize));
	const paginationPages = $derived.by(() => {
		const pageCount = Math.min(5, totalPages);
		const firstPage = Math.min(Math.max(data.page - 3, 1), Math.max(totalPages - pageCount + 1, 1));

		return Array.from({ length: pageCount }, (_, index) => firstPage + index);
	});

	const pageHref = (targetPage: number) => {
		const searchParams = new URLSearchParams(page.url.searchParams);
		searchParams.set('page', String(targetPage));
		return `${page.url.pathname}?${searchParams.toString()}${page.url.hash}`;
	};
</script>

<svelte:head>
	<title>Gallery | Open Artwork</title>
	<meta
		name="description"
		content="Explore thousands of public domain artworks from museum collections around the world. Discover artists, art movements, historical works, and inspiration for your next creative project."
	/>
</svelte:head>

<section
	id="gallery"
	class="flex min-h-screen flex-col px-6 py-15 sm:px-10 md:px-15 md:py-20 lg:gap-20 xl:p-25 2xl:p-30"
>
	<div class="group flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto" use:inView>
		<h1
			class="translate-y-5 text-5xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 md:text-6xl"
		>
			Search art
		</h1>

		<p
			class="translate-y-5 opacity-0 transition-all delay-160 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 lg:w-1/2"
		>
			Explore thousands of public domain artworks from museums and collections around the world.
			Explore different periods and mediums, and find inspiration for your next creative project.
		</p>
	</div>

	<p class="self-end">{data.totalResults} results</p>
	<div class="grid grid-cols-1 gap-15 md:grid-cols-2 lg:grid-cols-3 lg:gap-20 xl:grid-cols-4">
		{#each data.artworks as preview}
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

		<div class="col-span-1 mx-auto flex flex-col gap-5.5 md:col-span-2 lg:col-span-3 xl:col-span-4">
			<nav class="flex items-center gap-2.5" aria-label="Artwork pages">
				{#if data.page > 1}
					<Link variant="primary" href={pageHref(data.page - 1)} aria-label="Previous page">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="32"
							height="32"
							class="size-6.5 fill-bone-50"
							viewBox="0 0 256 256"
							aria-hidden="true"
							><path
								d="M165.66,202.34a8,8,0,0,1-11.32,11.32l-80-80a8,8,0,0,1,0-11.32l80-80a8,8,0,0,1,11.32,11.32L91.31,128Z"
							></path></svg
						>
					</Link>
				{:else}
					<Button variant="primary" disabled aria-label="Previous page">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="32"
							height="32"
							class="size-6.5 fill-bone-50"
							viewBox="0 0 256 256"
							aria-hidden="true"
							><path
								d="M165.66,202.34a8,8,0,0,1-11.32,11.32l-80-80a8,8,0,0,1,0-11.32l80-80a8,8,0,0,1,11.32,11.32L91.31,128Z"
							></path></svg
						>
					</Button>
				{/if}

				{#each paginationPages as paginationPage}
					{#if paginationPage === data.page}
						<Button variant="outline" disabled aria-current="page">{paginationPage}</Button>
					{:else}
						<Link variant="secondary" href={pageHref(paginationPage)}>{paginationPage}</Link>
					{/if}
				{/each}

				{#if data.hasMore}
					<Link variant="primary" href={pageHref(data.page + 1)} aria-label="Next page">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="32"
							height="32"
							class="size-6.5 rotate-180 fill-bone-50"
							viewBox="0 0 256 256"
							aria-hidden="true"
							><path
								d="M165.66,202.34a8,8,0,0,1-11.32,11.32l-80-80a8,8,0,0,1,0-11.32l80-80a8,8,0,0,1,11.32,11.32L91.31,128Z"
							></path></svg
						>
					</Link>
				{:else}
					<Button variant="primary" disabled aria-label="Next page">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="32"
							height="32"
							class="size-6.5 rotate-180 fill-bone-50"
							viewBox="0 0 256 256"
							aria-hidden="true"
							><path
								d="M165.66,202.34a8,8,0,0,1-11.32,11.32l-80-80a8,8,0,0,1,0-11.32l80-80a8,8,0,0,1,11.32,11.32L91.31,128Z"
							></path></svg
						>
					</Button>
				{/if}
			</nav>

			<Link href="#gallery" variant="tertiary" cn="mx-auto">Back to top</Link>
		</div>
	</div>
</section>
