<script lang="ts">
	import * as Form from '$lib/components/ui/form/index';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zod4Client } from 'sveltekit-superforms/adapters';
	import { loginSchema, registerSchema, type RegisterSchema } from './schema';
	import { Input } from '$lib/components/ui/input';
	import { invalidateAll } from '$app/navigation';
	import { toast } from 'svelte-sonner';

	type Props = {
		form: SuperValidated<Infer<RegisterSchema>>;
		closeDialog: () => void;
	};

	const props: Props = $props();

	const form = superForm(props.form, {
		validators: zod4Client(registerSchema),
		onUpdate: async ({ result }) => {
			if (result.type === 'success') {
				await invalidateAll();
				props.closeDialog();
				toast.success('Account registered');
			} else if (result.type === 'failure') {
				if ('message' in result.data) toast.error(result.data.message);
			}
		}
	});

	const { form: formData, enhance } = form;
</script>

<form method="POST" action="/register" class="space-y-4" use:enhance>
	<Form.Field {form} name="email">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Email</Form.Label>
				<Input
					{...props}
					bind:value={$formData.email}
					placeholder="test@example.com"
					autocomplete="email"
				/>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="username">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Username</Form.Label>
				<Input
					{...props}
					bind:value={$formData.username}
					placeholder="Username"
					autocomplete="username"
				/>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="password">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Password</Form.Label>
				<Input
					{...props}
					bind:value={$formData.password}
					placeholder="●●●●●●●●"
					type="password"
					autocomplete="new-password"
				/>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="confirmPassword">
		<Form.Control>
			{#snippet children({ props })}
				<Form.Label>Confirm Password</Form.Label>
				<Input
					{...props}
					bind:value={$formData.confirmPassword}
					placeholder="●●●●●●●●"
					type="password"
					autocomplete="new-password"
				/>
			{/snippet}
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button class="w-full">Register</Form.Button>
</form>
