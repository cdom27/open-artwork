<script lang="ts">
	import { footerLinks } from '$lib/data/nav-links';
	import { inView } from '$lib/actions/in-view';

	const date = new Date();
	const year = date.getFullYear();
</script>

<footer
	class="group grid grid-cols-1 gap-15 bg-bone-100 px-3 py-15 text-charcoal-900 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:items-center"
	use:inView
>
	<div
		class="relative text-bone-100 opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100"
	>
		<span
			class="absolute inset-0 z-10 flex items-center justify-center text-3xl sm:text-5xl lg:text-4xl xl:text-5xl"
			>Open Artwork</span
		>

		<div class="absolute inset-0 bg-linear-to-t via-charcoal-900/20 to-charcoal-900/0"></div>

		<div class="flex size-full">
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

					{#each category.links as link, i}
						<li
							style="--i: {i + 1}"
							class="underline-offset-4 opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
						>
							<a
								href={link.href}
								class="translate-y-5 underline decoration-bone-100 decoration-2 transition-all duration-250 hover:decoration-charcoal-400 active:decoration-charcoal-600"
								>{link.name}
							</a>
						</li>
					{/each}
				</ul>
			{/each}
		</div>

		<div class="flex flex-col gap-5.5 sm:flex-row sm:justify-between">
			<p
				class="translate-y-5 opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
				style="--i: 5"
			>
				Made by
				<a
					href="https://swe.cd"
					target="_blank"
					rel="noopener noreferrer"
					class="underline decoration-bone-100 decoration-2 underline-offset-4 transition-all duration-250 hover:decoration-charcoal-400 active:decoration-charcoal-600"
					>swe.cd</a
				>
			</p>

			<p
				style="--i: 5"
				class="translate-y-5 opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				© {year} Open Artwork
			</p>
		</div>
	</div>
</footer>
