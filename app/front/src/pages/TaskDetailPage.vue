<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const taskId = computed(() => route.params.id)
</script>

<template>
  <section class="task-detail">
    <div class="task-panel">
      <div class="panel-header">
        <div>
          <p class="eyebrow">Task details</p>
          <h2>Task #{{ taskId }}</h2>
          <p class="muted">Review context, ownership, and progress.</p>
        </div>
        <span class="pill">In progress</span>
      </div>
      <div class="panel-body">
        <div class="detail-row">
          <span class="label">Owner</span>
          <span>Imani R.</span>
        </div>
        <div class="detail-row">
          <span class="label">Due date</span>
          <span>Feb 18, 2026</span>
        </div>
        <div class="detail-row">
          <span class="label">Priority</span>
          <span>High</span>
        </div>
        <div class="detail-row">
          <span class="label">Description</span>
          <span>
            Align mobile layout updates with the latest customer feedback and
            ensure handoff assets are ready for engineering.
          </span>
        </div>
      </div>
      <div class="panel-actions">
        <button class="btn ghost">Add comment</button>
        <button class="btn primary">Mark as done</button>
      </div>
    </div>
  </section>
</template>
