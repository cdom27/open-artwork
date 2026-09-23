<script lang="ts">
	import { footerLinks } from '$lib/data/nav-links';
</script>

<footer class="grid grid-cols-1 gap-15 bg-black p-3 py-15 text-white">
	<div class="h-40 w-full rounded-sm bg-fuchsia-500"></div>

	<div class="grid grid-cols-1 gap-5">
		{#each footerLinks as category}
			<ul class="flex flex-col gap-0.5">
				<li>{category.name}</li>

				{#each category.links as link}
					<li><a href={link.href}>{link.name}</a></li>
				{/each}
			</ul>
		{/each}

		<div class="flex flex-col gap-5">
			<p>
				Made by <a
					href="https://swe.cd"
					target="_blank"
					rel="noopener noreferrer"
					class="underline underline-offset-4">swe.cd</a
				>
			</p>

			<p>© 2026 Open Artwork</p>
		</div>
	</div>
</footer>
