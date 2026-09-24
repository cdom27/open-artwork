<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLAnchorAttributes } from 'svelte/elements';
	import { variants } from '$lib/utils/link-button-variants';

	interface AnchorProps extends HTMLAnchorAttributes {
		variant: 'primary' | 'secondary' | 'tertiary';
		children: Snippet<[]>;
		cn?: string;
	}

	let { variant = 'primary', children, cn = '', ...rest }: AnchorProps = $props();

	const variantClass = $derived(variants[variant]);
	const classes = $derived(
		`${variantClass} transition-all duration-300 cursor-pointer rounded-sm py-1.5 px-2.5 hover:rounded-lg active:rounded-xl self-start ${cn}`.trim()
	);
</script>

<a class={classes} {...rest}>{@render children()}</a>
