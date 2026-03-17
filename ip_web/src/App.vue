<script setup lang="ts">
import { ref, computed } from 'vue'

const ip = ref('82.158.225.153')
const channels = ref<string[]>(['A', 'B', 'C', 'D'])
const allChannels: Record<string, string> = { 'A': '超时', 'B': '无法访问', 'C': 'ip-api.com', 'D': 'ipinfo.io' }
const touched = ref(false)
const loading = ref(false)
const result = ref<Record<string, any> | null>(null)
const errorMsg = ref('')

const ipPattern = /^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$/

const ipError = computed(() => {
  if (!touched.value || ip.value === '') return ''
  if (!ipPattern.test(ip.value)) return 'IP 格式不正确，请输入合法的 IPv4 地址'
  return ''
})

const canQuery = computed(() => {
  return ipPattern.test(ip.value) && channels.value.length > 0
})

async function query() {
  touched.value = true
  if (!canQuery.value) return
  loading.value = true
  result.value = null
  errorMsg.value = ''
  try {
    const params = new URLSearchParams()
    params.set('ip', ip.value)
    channels.value.forEach(ch => params.append('channels', ch))
    const res = await fetch(`http://localhost:9525/query?${params.toString()}`)
    if (!res.ok) {
      errorMsg.value = `请求失败: ${res.status} ${res.statusText}`
      return
    }
    result.value = await res.json()
  } catch (e: any) {
    errorMsg.value = `请求异常: ${e.message}`
  } finally {
    loading.value = false
  }
}

function onInput() {
  if (!touched.value && ip.value.length > 0) touched.value = true
}
</script>

<template>
  <div class="page">
    <div class="card">
      <div class="search-row">
        <div class="input-wrap" :class="{ 'has-error': ipError }">
          <input
            v-model="ip"
            type="text"
            placeholder="请输入 IP 地址，例如 8.8.8.8"
            @input="onInput"
            @keyup.enter="query"
          />
          <button :disabled="!canQuery || loading" @click="query">
            {{ loading ? '查询中...' : '查询' }}
          </button>
        </div>
        <p v-if="ipError" class="error-text">{{ ipError }}</p>
      </div>

      <div class="channels">
        <label v-for="(desc, ch) in allChannels" :key="ch" class="channel-item">
          <input type="checkbox" :value="ch" v-model="channels" />
          <span class="checkmark"></span>
          <span class="label-text">渠道{{ ch }}({{ desc }})</span>
        </label>
      </div>

      <div v-if="errorMsg" class="result-error">{{ errorMsg }}</div>

      <div v-if="result" class="result-box">
        <pre>{{ JSON.stringify(result, null, 2) }}</pre>
      </div>
    </div>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background: #f0f4f8;
  min-height: 100vh;
}
</style>

<style scoped>
.page {
  min-height: 100vh;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding: 20px;
  padding-top: 10vh;
}

.card {
  width: 100%;
  max-width: 50%;
  padding: 40px 36px 36px;
}

.title {
  text-align: center;
  font-size: 22px;
  font-weight: 600;
  color: #2d3748;
  margin-bottom: 28px;
}

.search-row {
  margin-bottom: 24px;
}

.input-wrap {
  display: flex;
  border: 2px solid #e2e8f0;
  border-radius: 48px;
  overflow: hidden;
  transition: border-color 0.2s;
}
.input-wrap:focus-within {
  border-color: #38b2ac;
}
.input-wrap.has-error {
  border-color: #e53e3e;
}

.input-wrap input[type="text"] {
  flex: 1;
  border: none;
  outline: none;
  padding: 14px 20px;
  font-size: 15px;
  color: #2d3748;
  background: transparent;
  min-width: 0;
}
.input-wrap input[type="text"]::placeholder {
  color: #a0aec0;
}

.input-wrap button {
  border: none;
  background: linear-gradient(135deg, #38b2ac, #319795);
  color: #fff;
  padding: 14px 32px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  transition: opacity 0.2s;
  letter-spacing: 2px;
}
.input-wrap button:hover:not(:disabled) {
  opacity: 0.9;
}
.input-wrap button:disabled {
  background: #a0aec0;
  cursor: not-allowed;
}

.error-text {
  color: #e53e3e;
  font-size: 13px;
  margin-top: 8px;
  padding-left: 20px;
}

.channels {
  display: flex;
  justify-content: center;
  gap: 28px;
  flex-wrap: wrap;
}

.channel-item {
  display: flex;
  align-items: center;
  cursor: pointer;
  margin-right:40px;
  user-select: none;
  position: relative;
}

.channel-item input[type="checkbox"] {
  position: absolute;
  opacity: 0;
  width: 0;
  height: 0;
}

.checkmark {
  width: 20px;
  height: 20px;
  border: 2px solid #38b2ac;
  border-radius: 4px;
  margin-right: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background 0.15s, border-color 0.15s;
  flex-shrink: 0;
}
.checkmark::after {
  content: '';
  display: none;
  width: 5px;
  height: 10px;
  border: solid #fff;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
  margin-bottom: 2px;
}
.channel-item input:checked ~ .checkmark {
  background: #38b2ac;
  border-color: #38b2ac;
}
.channel-item input:checked ~ .checkmark::after {
  display: block;
}

.label-text {
  font-size: 15px;
  color: #4a5568;
  font-weight: 500;
}

.result-error {
  margin-top: 24px;
  padding: 12px 16px;
  background: #fff5f5;
  border: 1px solid #feb2b2;
  border-radius: 8px;
  color: #c53030;
  font-size: 14px;
  text-align: center;
}

.result-box {
  margin-top: 24px;
  padding: 16px;
}
.result-box pre {
  font-size: 13px;
  color: #2d3748;
  white-space: pre-wrap;
  word-break: break-all;
  font-family: 'SF Mono', Monaco, Menlo, Consolas, monospace;
}
</style>
