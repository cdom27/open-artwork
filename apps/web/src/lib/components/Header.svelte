<script lang="ts">
	import { headerLinks } from '$lib/data/nav-links';
	import { inView } from '$lib/actions/in-view';
	import Button from './Button.svelte';
	import Link from './Link.svelte';

	let open = $state(true);
</script>

<header class="group relative p-2 sm:px-10 sm:py-5 md:px-15">
	<div class="flex items-center justify-between">
		<a href="/" onclick={() => (open = false)}>Open Artwork</a>

		<div class="bg-bone-50">
			<Button
				variant={open ? 'outline' : 'tertiary'}
				onclick={() => (open = !open)}
				cn="md:hidden border border-bone-50">{open ? 'Close Menu' : 'Menu'}</Button
			>
		</div>

		<nav class="hidden md:flex">
			<ul class="flex gap-5.5">
				{#each headerLinks as linkGroup}
					<Link variant="tertiary" href={linkGroup.href} onclick={() => (open = false)}
						>{linkGroup.name}</Link
					>
				{/each}
			</ul>
		</nav>
	</div>

	<nav class={`absolute inset-x-2 mt-5.5 sm:inset-x-10 md:inset-x-15 ${open ? 'flex' : 'hidden'}`}>
		<ul class=" ml-auto flex flex-col gap-5.5">
			{#each headerLinks as linkGroup}
				<li class="flex self-end">
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
</header>
