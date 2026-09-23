<script lang="ts">
	import { footerLinks } from '$lib/data/nav-links';
</script>

<footer
	class="grid grid-cols-1 gap-15 bg-bone-100 px-3 py-15 text-charcoal-900 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:items-center"
>
	<div class="relative text-bone-100">
		<span
			class="absolute inset-0 z-10 flex items-center justify-center text-3xl sm:text-5xl lg:text-4xl xl:text-5xl"
			>Open Artwork</span
		>

		<div class="absolute inset-0 bg-linear-to-t via-charcoal-900/20 to-charcoal-900/0"></div>

		<div class="flex size-full bg-red-100">
			<img
				src="/images/repose-john-white-alexander-1895.webp"
				class="rounded-sm object-cover"
				alt="Repose by John White Alexander, 1895"
			/>
		</div>
	</div>

	<div class="grid grid-cols-1 gap-5.5 sm:gap-15">
		<div class="flex flex-col gap-5.5 sm:flex-row sm:justify-between">
			{#each footerLinks as category}
				<ul class="flex flex-col gap-1.5">
					<li class="text-charcoal-600">{category.name}</li>

					{#each category.links as link}
						<li>
							<a
								href={link.href}
								class="underline decoration-bone-100 decoration-2 underline-offset-4 transition-all duration-250 hover:decoration-charcoal-400"
								>{link.name}</a
							>
						</li>
					{/each}
				</ul>
			{/each}
		</div>

		<div class="flex flex-col gap-5.5 sm:flex-row sm:justify-between">
			<p>
				Made by <a
					href="https://swe.cd"
					target="_blank"
					rel="noopener noreferrer"
					class="underline decoration-bone-100 decoration-2 underline-offset-4 transition-all duration-250 hover:decoration-charcoal-400"
					>swe.cd</a
				>
			</p>

			<p>© 2026 Open Artwork</p>
		</div>
	</div>
</footer>
