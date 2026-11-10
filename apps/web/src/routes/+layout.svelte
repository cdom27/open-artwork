<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Footer from '$lib/components/Footer.svelte';
	import Header from '$lib/components/Header.svelte';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<aside
	class="absolute inset-x-0 top-0 z-20 bg-charcoal-900 px-4 py-2.5 text-center text-sm leading-5 text-bone-50 sm:px-10"
>
	This live version is a prototype showcasing the gallery and color extractions from artworks
	processed so far. New data and fixes will arrive as they are processed. Discovery views, image
	processing tools, advanced palette and color tools, and the developer API are not enabled yet.
</aside>

<Header />

<main class="min-h-screen scroll-smooth">{@render children()}</main>

<Footer />
