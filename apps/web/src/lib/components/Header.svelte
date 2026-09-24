<script lang="ts">
	import { fly } from 'svelte/transition';
	import { headerLinks } from '$lib/data/nav-links';
	import Button from './Button.svelte';
	import Link from './Link.svelte';
	import Logo from '$lib/assets/favicon.svg';

	let open = $state(false);
</script>

<header class="relative p-2 sm:px-10 sm:py-5 md:px-15">
	<div class="flex items-center justify-between">
		<a
			href="/"
			onclick={() => (open = false)}
			class="translate-color flex items-center gap-2 decoration-bone-50 decoration-2 underline-offset-4 duration-250 hover:underline hover:decoration-charcoal-400 active:decoration-charcoal-600"
			><img src={Logo} alt="Open Artwork" /><span class="text-xl">Open Artwork</span></a
		>

		<div class="bg-bone-50">
			<Button
				variant={open ? 'outline' : 'tertiary'}
				onclick={() => (open = !open)}
				cn="md:hidden border border-bone-50">{open ? 'Close Menu' : 'Menu'}</Button
			>
		</div>

		<nav class="hidden md:flex">
			<ul class="flex gap-5.5">
				{#each headerLinks as link}
					<Link variant="tertiary" href={link.href} onclick={() => (open = false)}>{link.name}</Link
					>
				{/each}
			</ul>
		</nav>
	</div>

	{#if open}
		<nav
			class="absolute inset-x-2 mt-5.5 flex sm:inset-x-10 md:inset-x-15"
			transition:fly={{ y: -10, duration: 300 }}
		>
			<ul class="ml-auto flex flex-col gap-5.5">
				{#each headerLinks as linkGroup, i}
					<li
						class="flex self-end"
						in:fly|global={{ y: -10, duration: 300, delay: i * 80 }}
						out:fly|global={{ y: -10, duration: 200, delay: (headerLinks.length - 1 - i) * 60 }}
					>
						<Link
							variant="tertiary"
							cn="self-end"
							href={linkGroup.href}
							onclick={() => (open = false)}>{linkGroup.name}</Link
						>
					</li>
				{/each}
			</ul>
		</nav>
	{/if}
</header>
