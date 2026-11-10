<script lang="ts">
	import { inView } from '$lib/actions/in-view';
	import CheckIcon from '$lib/components/icons/CheckIcon.svelte';
	import CopyIcon from '$lib/components/icons/CopyIcon.svelte';
	import Figure from '$lib/components/Figure.svelte';
	import { getAccessiblePairings } from '$lib/utils/accessible-pairings';
	import { getDisplayHex, getTextColor, handleCopyColor } from '$lib/utils/colors';
	import { artworkMoods, type ArtworkMood } from '$lib/types/artwork';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	let copiedHex = $state<string | null>(null);
	let moodOpen = $state(false);
	let constituentsOpen = $state(false);

	const artist = $derived(
		data.creditLine.length > 95 ? `${data.creditLine.slice(0, 95)}...` : data.creditLine
	);
	const title = $derived(data.title.length > 95 ? `${data.title.slice(0, 95)}...` : data.title);
	let selectedMood = $state<ArtworkMood>('Colorful');
	const selectedPalette = $derived(
		data.palettes.find((palette) => palette.mood === selectedMood) ?? data.palettes[0]
	);
	const primaryConstituent = $derived(data.constituents[0] ?? 'Unknown artist');
	const additionalConstituents = $derived(Math.max(data.constituents.length - 1, 0));
	const informationItems = $derived([
		{ label: '', value: data.dateDisplay },
		{ label: 'Medium', value: data.medium || 'Unknown' },
		{ label: 'Classification', value: data.classification || 'Unknown' },
		{ label: 'Department', value: data.department || 'Unknown' },
		{ label: 'Credit', value: data.creditLine || 'Unknown' },
		...(data.tags.length > 0 ? [{ label: 'Related to', value: data.tags.join(', ') }] : []),
		{ label: 'Dimensions', value: data.dimensions || 'Unknown' }
	]);
	const accessiblePairings = $derived(
		selectedPalette ? getAccessiblePairings(selectedPalette.colors.map((color) => color.hex)) : []
	);

	const showCopiedState = (hex: string) => {
		copiedHex = hex;
		window.setTimeout(() => {
			if (copiedHex === hex) copiedHex = null;
		}, 1500);
	};
</script>

<svelte:head>
	<title>{data.title}, {data.creditLine}, {data.dateDisplay} | Open Artwork</title>
	<meta
		name="description"
		content={`See the color palette and open data of: ${data.title}, ${data.creditLine}, ${data.dateDisplay}.`}
	/>
</svelte:head>

<section
	class="group flex flex-col px-6 py-15 sm:px-10 md:px-15 md:py-20 xl:p-25 2xl:p-30"
	use:inView
>
	<h1 class="sr-only">
		{data.title} by {data.creditLine}, {data.dateDisplay}
	</h1>

	<Figure
		cn="mx-auto w-full opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100 sm:w-1/2"
		img={{
			src: data.fullUrl,
			width: 1023,
			height: 724,
			loading: 'eager',
			fetchpriority: 'high'
		}}
		artwork={{
			title,
			artist,
			date: data.dateDisplay
		}}
	/>
</section>

<section
	class="grid grid-cols-1 gap-5.5 px-6 py-15 sm:gap-7.5 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:gap-20 xl:p-25 2xl:p-30"
>
	<div class="group flex flex-col gap-7.5" use:inView>
		<div class="flex flex-col gap-7.5 sm:flex-row sm:items-center sm:justify-between sm:gap-0">
			<h2
				class="translate-y-5 text-3xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 lg:text-4xl"
			>
				Color Palette
			</h2>

			<div class="flex items-center gap-2.5 text-sm">
				<span>Mood</span>
				<details
					class="group relative text-charcoal-900"
					bind:open={moodOpen}
					onmouseleave={() => (moodOpen = false)}
				>
					<summary
						class="flex cursor-pointer list-none items-center justify-between gap-2 rounded-sm bg-bone-200 px-2.5 py-1.5 [&::-webkit-details-marker]:hidden"
					>
						<span>{selectedMood}</span>
						<svg
							class="size-5.5 transition-transform duration-200 group-open:rotate-180"
							viewBox="0 0 256 256"
							aria-hidden="true"
						>
							<path
								d="M208.49,80.49l-72,72a8,8,0,0,1-11.32,0l-72-72a8,8,0,0,1,11.32-11.32L128,132.69l63.17-63.52a8,8,0,0,1,11.32,11.32Z"
							></path>
						</svg>
					</summary>

					{#if moodOpen}
						<div class="absolute right-0 z-10 min-w-full pt-2.5">
							<ul
								class="overflow-hidden rounded-sm bg-bone-100 py-1 shadow-xs"
								aria-label="Palette moods"
							>
								{#each artworkMoods as mood}
									<li>
										<button
											type="button"
											class={`block w-full px-3 py-2 text-left text-sm transition-colors hover:bg-bone-200 active:bg-bone-300 ${mood === selectedMood ? 'bg-charcoal-100 font-medium' : ''}`}
											onclick={() => {
												selectedMood = mood;
												moodOpen = false;
											}}
										>
											{mood}
										</button>
									</li>
								{/each}
							</ul>
						</div>
					{/if}
				</details>
			</div>
		</div>

		{#if selectedPalette}
			<div class="flex h-96 flex-col overflow-hidden rounded-sm sm:h-96 sm:flex-row">
				{#each selectedPalette.colors as color, i}
					<button
						type="button"
						style={`--i: ${i + 1}; background-color: ${color.hex}`}
						class={`group/swatch relative flex min-h-0 flex-1 translate-y-5 items-center justify-center opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 hover:cursor-pointer sm:min-h-20 ${getTextColor(color.hex)}`}
						onclick={() => handleCopyColor(color.hex, () => showCopiedState(color.hex))}
						aria-label={`Copy ${getDisplayHex(color.hex)} to clipboard`}
					>
						<span
							class="absolute inset-0 flex items-center justify-center text-center text-sm sm:inset-x-0 sm:inset-y-auto sm:bottom-5"
						>
							{#if copiedHex === color.hex}
								<span class="flex items-center gap-1 sm:hidden">
									Copied <CheckIcon cn="size-4" />
								</span>
								<span class="hidden sm:inline">{getDisplayHex(color.hex)}</span>
							{:else}
								{getDisplayHex(color.hex)}
							{/if}
						</span>
						{#if copiedHex === color.hex}
							<span
								class="pointer-events-none absolute inset-x-0 bottom-[calc(50%+1.5rem)] hidden justify-center sm:bottom-14 sm:flex"
								aria-label="Copied"
							>
								<CheckIcon cn="size-5" />
							</span>
						{:else}
							<span
								class="pointer-events-none absolute inset-x-0 bottom-[calc(50%+1.5rem)] flex justify-center opacity-0 transition-opacity group-hover/swatch:opacity-100 sm:bottom-14"
								aria-hidden="true"
							>
								<CopyIcon cn="size-4" />
							</span>
						{/if}
					</button>
				{/each}
			</div>
		{/if}
	</div>

	<div class="group flex flex-col gap-7.5" use:inView>
		<h2
			class="translate-y-5 text-3xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 lg:text-4xl"
		>
			Accessible Pairings
		</h2>

		{#if accessiblePairings.length > 0}
			<div class="flex flex-col gap-7.5">
				{#each accessiblePairings as pairing, i}
					<div
						style="--i: {i + 1}"
						class="flex translate-y-5 items-center gap-3.5 opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
					>
						<div
							class="flex h-12 w-17 shrink-0 items-center justify-center rounded-sm"
							style={`background-color: ${pairing.background}; color: ${pairing.foreground}`}
						>
							<span class="text-xl">Aa</span>
						</div>
						<div class="flex flex-col">
							<p class="text-sm">
								#{getDisplayHex(pairing.foreground)} on #{getDisplayHex(pairing.background)}
							</p>
							<p class="text-sm text-charcoal-600">{pairing.ratio.toFixed(2)}:1</p>
						</div>
					</div>
				{/each}
			</div>
		{:else}
			<p class="text-charcoal-600">No WCAG AA pairings found in this palette.</p>
		{/if}
	</div>
</section>

<section
	class="group grid grid-cols-1 px-6 py-15 sm:px-10 md:px-15 md:py-20 lg:grid-cols-2 lg:gap-20 xl:p-25 2xl:p-30"
	use:inView
>
	<div class="flex flex-col gap-5.5 sm:gap-7.5 lg:my-auto">
		<h2
			class="translate-y-5 text-3xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 lg:text-4xl"
		>
			{data.title}
		</h2>

		<Figure
			showCaption="never"
			cn="lg:hidden opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100"
			img={{
				src: data.fullUrl,
				width: 1023,
				height: 724,
				loading: 'eager',
				fetchpriority: 'high'
			}}
			artwork={{
				title: data.title,
				artist,
				date: data.dateDisplay
			}}
		/>

		<div class="flex flex-col gap-1.5 pt-5.5">
			<p
				style="--i: 1"
				class="translate-y-5 opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
			>
				{primaryConstituent}
				{#if additionalConstituents > 0}
					<button
						type="button"
						class="cursor-pointer underline decoration-bone-100 decoration-2 underline-offset-4 transition-all duration-250 hover:decoration-charcoal-400 active:decoration-charcoal-600"
						aria-expanded={constituentsOpen}
						onclick={() => (constituentsOpen = !constituentsOpen)}
					>
						+{additionalConstituents}
					</button>
				{/if}
			</p>

			{#if constituentsOpen}
				{#each data.constituents.slice(1) as constituent, i}
					<p
						style={`--i: ${i + 2}`}
						class="translate-y-5 opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
					>
						{constituent}
					</p>
				{/each}
			{/if}

			{#each informationItems as item, i}
				<p
					style={`--i: ${i + 2 + (constituentsOpen ? additionalConstituents : 0)}`}
					class="translate-y-5 opacity-0 transition-all delay-[calc(var(--i)*80ms)] duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100"
				>
					{#if item.label}<span class="text-charcoal-600">{item.label} |</span>{/if}
					{item.value}
				</p>
			{/each}
		</div>
	</div>

	<Figure
		cn="hidden lg:block opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100"
		showCaption="never"
		img={{
			src: data.fullUrl,
			width: 1023,
			height: 724,
			loading: 'lazy',
			fetchpriority: 'high'
		}}
		artwork={{
			title: data.title,
			artist,
			date: data.dateDisplay
		}}
	/>
</section>

{#if data.relatedArtworks.length > 0}
	<section class="flex flex-col gap-7.5 px-6 py-15 sm:px-10 md:px-15 md:py-20 xl:p-25 2xl:p-30">
		<div class="group" use:inView>
			<h2
				class="translate-y-5 text-3xl font-normal opacity-0 transition-all delay-80 duration-400 group-[&.in-view]:translate-y-0 group-[&.in-view]:opacity-100 lg:text-4xl"
			>
				Related artworks
			</h2>
		</div>

		<div class="grid grid-cols-1 gap-15 sm:grid-cols-2 lg:grid-cols-3 lg:gap-20 xl:grid-cols-5">
			{#each data.relatedArtworks as artwork}
				<a href={`/artwork/${artwork.slug}`} use:inView class="group">
					<Figure
						cn="aspect-[3/2] w-full opacity-0 transition-all duration-400 group-[&.in-view]:opacity-100 group-hover:scale-102"
						artwork={{
							title: artwork.title,
							artist: artwork.primaryArtistName,
							date: artwork.dateDisplay
						}}
						img={{ src: artwork.thumbnailUrl }}
						showCaption="never"
					/>
				</a>
			{/each}
		</div>
	</section>
{/if}
