<script lang="ts">
	import { headerLinks } from '$lib/data/nav-links';
	import { inView } from '$lib/actions/in-view';

	let open = $state(false);
</script>

<header class="group relative p-2 sm:px-10 sm:py-5 md:px-15">
	<div class="flex items-center justify-between">
		<a href="/">Open Artwork</a>

		<button class="cursor-pointer md:hidden" onclick={() => (open = !open)}>Menu</button>

		<nav class="hidden md:flex">
			<ul class="flex gap-5.5">
				{#each headerLinks as linkGroup}
					<li>
						<a href={linkGroup.href}>{linkGroup.name}</a>
					</li>
				{/each}
			</ul>
		</nav>
	</div>

	<nav
		class={`absolute inset-x-2 mt-2 sm:inset-x-10 sm:mt-5 md:inset-x-15 ${open ? 'flex' : 'hidden'}`}
	>
		<ul class="ml-auto flex flex-col gap-2.5">
			{#each headerLinks as linkGroup}
				<li class="self-end">
					<a href={linkGroup.href} class="self-end" onclick={() => (open = false)}
						>{linkGroup.name}</a
					>
				</li>
			{/each}

			<li class="self-end">
				<button class="cursor-pointer md:hidden" onclick={() => (open = false)}>Close Menu</button>
			</li>
		</ul>
	</nav>
</header>
