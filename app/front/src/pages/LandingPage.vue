<script setup>
const highlights = [
  {
    title: 'Plan your work',
    description: 'Organize tasks into clear stages and keep everyone aligned.',
  },
  {
    title: 'Track progress',
    description: 'Visual boards and task details keep momentum transparent.',
  },
  {
    title: 'Stay on schedule',
    description: 'Smart reminders and ownership reduce dropped work.',
  },
]
</script>

<template>
  <div>
    <section class="hero">
        <div class="hero-content">
        <p class="eyebrow">Taskflow</p>
        <h1>Make every task visible and actionable.</h1>
        <p class="lead">
            Taskflow helps teams move faster with a simple, focused workflow for
            managing tasks and priorities.
        </p>
        <div class="hero-actions">
            <RouterLink class="btn primary" to="/login">Get started</RouterLink>
            <RouterLink class="btn ghost" to="/dashboard">View demo board</RouterLink>
        </div>
        </div>
        <div class="hero-card">
        <div class="card-header">
            <div>
            <p class="card-title">Today</p>
            <p class="card-subtitle">3 tasks need attention</p>
            </div>
            <span class="pill">Live</span>
        </div>
        <div class="card-body">
            <div class="mini-task">
            <span class="dot urgent"></span>
            Review onboarding checklist
            </div>
            <div class="mini-task">
            <span class="dot"></span>
            Update mobile flow specs
            </div>
            <div class="mini-task">
            <span class="dot"></span>
            Prepare sprint demo deck
            </div>
        </div>
        </div>
    </section>

    <section class="features">
        <div v-for="item in highlights" :key="item.title" class="feature-card">
        <h3>{{ item.title }}</h3>
        <p>{{ item.description }}</p>
        </div>
    </section>
  </div>
</template>
