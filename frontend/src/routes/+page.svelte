<script lang="ts">
	import type { IShortenURLResponse } from '../lib';

	let url = $state('');
	let shortenedPath = $state<string | null>(null);

	const shortenURL = async () => {
		const res = await fetch('http://localhost:8080/u', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ url })
		});

		if (res.ok) {
			const { urlPath } = (await res.json()) as IShortenURLResponse;
			shortenedPath = urlPath;
		}
	};
</script>

<main>
	<h1>Shorten</h1>
	<input type="text" placeholder="URL:" bind:value={url} />

	<button onclick={shortenURL}>Shorten</button>

	{#if shortenedPath}
		<p>Host: {window.location.host}</p>
		<p>Shortened path: {shortenedPath}</p>
	{/if}
</main>
