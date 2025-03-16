<!-- LanguageSwitcher.vue -->
<template>
  <div class="language-switcher">
    <a-dropdown :trigger="['click']">
      <a class="ant-dropdown-link" @click.prevent>
        {{ getCurrentLanguageName }} <down-outlined />
      </a>
      <template #overlay>
        <a-menu @click="handleLanguageChange">
          <a-menu-item v-for="lang in languages" :key="lang.code">
            {{ lang.name }}
          </a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { DownOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import axios from 'axios'

export default {
  name: 'LanguageSwitcher',
  components: {
    DownOutlined
  },
  setup() {
    const languages = ref([])
    const currentLanguage = ref('zh-CN')
    
    // 获取支持的语言列表
    const fetchLanguages = async () => {
      try {
        const response = await axios.get('/api/common/getLanguages')
        if (response.data.code === 0) {
          languages.value = response.data.data
        }
      } catch (error) {
        console.error('获取语言列表失败:', error)
      }
    }
    
    // 获取当前语言
    const fetchCurrentLanguage = async () => {
      try {
        const response = await axios.get('/api/common/getCurrentLanguage')
        if (response.data.code === 0) {
          currentLanguage.value = response.data.data
        }
      } catch (error) {
        console.error('获取当前语言失败:', error)
      }
    }
    
    // 切换语言
    const handleLanguageChange = async (e) => {
      const langCode = e.key
      try {
        const response = await axios.get(`/api/common/setLanguage?lang=${langCode}`)
        if (response.data.code === 0) {
          currentLanguage.value = langCode
          message.success('语言切换成功，页面将在 1 秒后刷新')
          // 刷新页面以应用新的语言设置
          setTimeout(() => {
            window.location.reload()
          }, 1000)
        }
      } catch (error) {
        console.error('语言切换失败:', error)
        message.error('语言切换失败')
      }
    }
    
    // 获取当前语言名称
    const getCurrentLanguageName = computed(() => {
      const lang = languages.value.find(l => l.code === currentLanguage.value)
      return lang ? lang.name : '简体中文'
    })
    
    onMounted(() => {
      fetchLanguages()
      fetchCurrentLanguage()
    })
    
    return {
      languages,
      currentLanguage,
      handleLanguageChange,
      getCurrentLanguageName
    }
  }
}
</script>

<style scoped>
.language-switcher {
  display: inline-block;
  cursor: pointer;
}
</style> 