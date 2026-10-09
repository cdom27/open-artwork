<script lang="ts">
	import type { HTMLImgAttributes } from 'svelte/elements';

	interface FigureProps {
		img: HTMLImgAttributes;
		artwork: {
			title: string;
			artist: string;
			date: string;
		};
		cn?: string;
		showCaption?: 'always' | 'onHover';
	}

	let { img, artwork, cn = '', showCaption = 'always' }: FigureProps = $props();

	const classes = $derived(`relative self-start ${cn}`.trim());
</script>

<figure class={classes}>
	<img
		{...img}
		alt={artwork.title + ' by ' + artwork.artist + ', ' + artwork.date}
		class="size-full rounded-sm object-cover"
	/>
	<figcaption
		class={`absolute inset-x-0 bottom-0 z-10 flex min-h-24 items-end bg-linear-to-t from-charcoal-900/50 to-charcoal-900/0 p-3.5 font-normal text-bone-50 ${
			showCaption === 'onHover'
				? 'opacity-0 transition-opacity duration-400 group-hover:opacity-100'
				: ''
		}`}
	>
		{artwork.title}<br />
		{artwork.artist}<br />
		{artwork.date}
	</figcaption>
</figure>
