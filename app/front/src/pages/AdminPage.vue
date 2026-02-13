<script setup>
const roles = [
  {
    name: 'Administrators',
    description: 'Manage users, permissions, and workspace settings.',
  },
  {
    name: 'Team leads',
    description: 'Create projects, assign tasks, and review progress.',
  },
  {
    name: 'Members',
    description: 'Update tasks, add comments, and close work items.',
  },
]
</script>

<template>
  <section class="admin">
    <header class="admin-header">
      <div>
        <p class="eyebrow">Admin</p>
        <h2>Workspace controls</h2>
        <p class="muted">Manage access, policies, and service settings.</p>
      </div>
      <button class="btn primary">Invite user</button>
    </header>

    <div class="admin-grid">
      <div class="admin-card">
        <h3>Access policies</h3>
        <p class="muted">
          Enforce single sign-on, multi-factor authentication, and domain rules.
        </p>
        <button class="btn ghost">Configure</button>
      </div>
      <div class="admin-card">
        <h3>Audit log</h3>
        <p class="muted">Track logins, role changes, and key exports.</p>
        <button class="btn ghost">View log</button>
      </div>
      <div class="admin-card">
        <h3>Billing</h3>
        <p class="muted">Update plan, invoices, and cost centers.</p>
        <button class="btn ghost">Manage billing</button>
      </div>
    </div>

    <div class="role-list">
      <h3>Role definitions</h3>
      <div class="role-item" v-for="role in roles" :key="role.name">
        <div>
          <h4>{{ role.name }}</h4>
          <p class="muted">{{ role.description }}</p>
        </div>
        <button class="btn ghost">Edit</button>
      </div>
    </div>
  </section>
</template>
