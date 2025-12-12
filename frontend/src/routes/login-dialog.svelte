<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index';
	import * as Tabs from '$lib/components/ui/tabs/index';
	import type { ComponentProps } from 'svelte';
	import LoginForm from './login-form.svelte';
	import RegisterForm from './register-form.svelte';
	import { loginDialogState } from '$lib/states.svelte';

	type Props = {
		loginForm: ComponentProps<typeof LoginForm>['form'];
		registerForm: ComponentProps<typeof RegisterForm>['form'];
	};

	const { loginForm, registerForm, ...props }: Props = $props();
	let selectedTab = $state('sign-in');
</script>

<Dialog.Root open={loginDialogState.open} onOpenChange={(open) => (loginDialogState.open = open)}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{selectedTab === 'sign-in' ? 'Sign In' : 'Register'}</Dialog.Title>
			<Dialog.Description
				>{selectedTab === 'sign-in'
					? 'Login to your account'
					: 'Register a new account'}</Dialog.Description
			>
		</Dialog.Header>
		<Tabs.Root value={selectedTab} onValueChange={(val) => (selectedTab = val)} class="w-full">
			<Tabs.List class="w-full">
				<Tabs.Trigger value="sign-in">Sign In</Tabs.Trigger>
				<Tabs.Trigger value="register">Register</Tabs.Trigger>
			</Tabs.List>
			<Tabs.Content value="sign-in">
				<LoginForm form={loginForm} />
			</Tabs.Content>
			<Tabs.Content value="register">
				<RegisterForm form={registerForm} />
			</Tabs.Content>
		</Tabs.Root>
	</Dialog.Content>
</Dialog.Root>
