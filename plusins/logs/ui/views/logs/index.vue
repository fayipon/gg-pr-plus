<template>
  <div>
    <a-card title="系统日志">
      <a-table :columns="columns" :data-source="data" row-key="id" :pagination="false" />
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { defHttp } from '@/utils/http';

const data = ref([]);

const columns = [
  { title: 'ID', dataIndex: 'id' },
  { title: '用户ID', dataIndex: 'user_id' },
  { title: '操作', dataIndex: 'action' },
  { title: '详情', dataIndex: 'detail' },
  { title: 'IP 地址', dataIndex: 'ip' },
  { title: '时间', dataIndex: 'create_time' },
];

onMounted(async () => {
  const res = await defHttp.post({ url: '/logs/list', params: { page: 1, pageSize: 20 } });
  // 处理返回可能有 { data, total } 或 直接数组
  data.value = res.data ?? res;
});
</script>
