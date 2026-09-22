<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let navLinks = [
		{
			label: 'Gallery',
			href: '/gallery'
		},
		{
			label: 'API',
			href: '/docs'
		},
		{
			label: 'About',
			href: '/about'
		}
	];

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<nav class="flex items-center justify-between p-4">
	<a href="/" class="hover:text-blue-600 hover:underline">Open Artwork</a>
	<ul class="flex items-center gap-3">
		{#each navLinks as link}
			<li><a href={link.href} class="hover:text-blue-600 hover:underline">{link.label}</a></li>
		{/each}
	</ul>
</nav>

{@render children()}
