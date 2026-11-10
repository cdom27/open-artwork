<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Footer from '$lib/components/Footer.svelte';
	import Header from '$lib/components/Header.svelte';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<aside
	class="flex flex-col gap-2.5 bg-charcoal-900 px-4 py-2.5 text-center font-normal text-bone-50 sm:px-10"
>
	<p>
		This live version is a prototype showcasing the gallery and color extractions from artworks
		processed so far.
	</p>
	<p>
		New data and fixes will arrive as they are processed. Discovery views, image processing tools,
		advanced palette and color tools, and the developer API are not enabled yet.
	</p>
</aside>

<Header />

<main class="min-h-screen scroll-smooth">{@render children()}</main>

<Footer />
