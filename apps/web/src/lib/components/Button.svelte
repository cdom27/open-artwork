<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';
	import { variants } from '$lib/utils/link-button-variants';

	interface ButtonProps extends HTMLButtonAttributes {
		variant: 'primary' | 'secondary' | 'tertiary' | 'outline';
		children: Snippet<[]>;
		cn?: string;
	}

	let { variant = 'primary', children, cn = '', ...rest }: ButtonProps = $props();

	const variantClass = $derived(variants[variant]);
	const classes = $derived(
		`${variantClass} cursor-pointer rounded-sm px-2.5 py-1.5 transition-all duration-300 hover:rounded-lg active:rounded-xl disabled:pointer-events-none disabled:cursor-not-allowed self-start ${cn}`.trim()
	);
</script>

<button class={classes} {...rest}>{@render children()}</button>
