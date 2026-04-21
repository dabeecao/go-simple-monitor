document.addEventListener('alpine:init', () => {
    Alpine.data('appData', () => ({
        lang: localStorage.getItem('lang') || 'vi',
        translations: {
            vi: {
                mobile_warning: 'Mở trên máy tính để có trải nghiệm quản lý tốt nhất',
                login_subtitle: 'Hệ thống quản lý thông minh',
                account: 'Tài khoản',
                password: 'Mật khẩu',
                username_ph: 'Nhập tên đăng nhập',
                login_btn: 'Đăng nhập',
                
                tab_dashboard: 'Tổng quan',
                tab_monitor: 'Giám sát',
                tab_terminal: 'Terminal',
                tab_settings: 'Cấu hình',

                search_ph: 'Tìm PID, Tiến trình...',
                lang_toggle: 'Ngôn ngữ / Language',
                theme_toggle: 'Giao diện',
                logout: 'Đăng xuất',

                server_info: 'Thông tin Máy chủ',
                os: 'Hệ điều hành',
                internal_ip: 'IP Nội bộ',
                net_traffic: 'Tổng Lưu lượng Mạng',
                total_up: 'Tổng Upload',
                total_down: 'Tổng Download',
                
                process: 'Tiến trình',
                network: 'Mạng',
                service: 'Dịch vụ',
                connection: 'Kết nối',
                kill: 'Kết thúc (Kill)',
                restart_term: 'Khởi động lại Terminal',

                cron_title: 'Tự động hóa (Cron)',
                cron_select: '-- Chọn chu kỳ --',
                cron_min: 'Mỗi phút',
                cron_5min: 'Mỗi 5 phút',
                cron_hour: 'Mỗi giờ',
                cron_day: 'Mỗi 00:00 hàng ngày',
                cron_cmd: 'Lệnh thực thi...',
                cron_empty: 'Chưa có Cronjob nào',

                tg_title: 'Cảnh báo Telegram',
                tg_cond: 'Điều kiện cảnh báo',
                tg_cpu: 'CPU vượt ngưỡng',
                tg_ram: 'RAM vượt ngưỡng',
                tg_disk: 'Ổ cứng đầy',
                save_cfg: 'Lưu Cấu Hình',

                cancel: 'Hủy',
                confirm: 'Xác nhận',

                msg_login_ok: 'Đăng nhập thành công',
                msg_conn_err: 'Lỗi kết nối server',
                msg_logout_title: 'Đăng xuất',
                msg_logout_desc: 'Bạn có chắc chắn muốn đăng xuất khỏi hệ thống?',
                msg_cron_add_err: 'Vui lòng điền đủ chu kỳ và lệnh!',
                msg_cron_added: 'Đã thêm Cronjob',
                msg_cron_del_title: 'Xóa Cronjob',
                msg_cron_del_desc: 'Bạn chắc chắn muốn xóa tác vụ này?',
                msg_deleted: 'Đã xóa',
                msg_cfg_saved: 'Đã lưu cấu hình',
                msg_cfg_err: 'Lỗi khi lưu cấu hình',
                msg_kill_title: 'Kết thúc tiến trình',
                msg_kill_desc: 'Ép buộc dừng tiến trình {pid} ({name})?',
                msg_kill_ok: 'Đã kill PID: {pid}',
                msg_kill_err: 'Lỗi khi kill tiến trình',
                msg_term_title: 'Khởi động lại Terminal?',
                msg_term_desc: 'Hành động này sẽ ngắt kết nối hiện tại và tạo một phiên Terminal mới. Bạn có chắc chắn không?',
                msg_term_ok: 'Đã làm mới phiên Terminal',
                cpu_load: 'Tải CPU',
                ram_usage: 'Sử dụng RAM',
                disk_usage: 'Dung lượng Đĩa',
                clock_speed: 'Xung nhịp',
                cores: 'Lõi',
                threads: 'Luồng',
                active: 'Hoạt động',
                cached: 'Bộ nhớ đệm',
                buffers: 'Bộ đệm dữ liệu',
                ports: 'Cổng kết nối',
                port: 'Cổng',
                bind_ip: 'IP Liên kết',
                cron_15min: 'Mỗi 15 phút',
                cron_30min: 'Mỗi 30 phút',
                cron_12hour: 'Mỗi 12 giờ',
                cron_sunday: 'Mỗi Chủ Nhật',
                cron_month: 'Ngày 1 hàng tháng',
                tg_test_empty: "Vui lòng nhập Bot Token và Chat ID trước khi test!",
                tg_test_sending: "Đang gửi tin nhắn test...",
                tg_test_success: "Gửi tin nhắn test thành công!",
                tg_test_failed: "Không thể gửi tin nhắn",
                error: "Lỗi",
                network_error: "Lỗi mạng khi kết nối đến máy chủ.",
                test_cfg: "Thử"
            },
            en: {
                mobile_warning: 'Open on desktop for the best management experience',
                login_subtitle: 'Smart Management System',
                account: 'Account',
                password: 'Password',
                username_ph: 'Enter username',
                login_btn: 'Login',
                
                tab_dashboard: 'Dashboard',
                tab_monitor: 'Monitor',
                tab_terminal: 'Terminal',
                tab_settings: 'Settings',

                search_ph: 'Search PID, Processes...',
                lang_toggle: 'Language / Ngôn ngữ',
                theme_toggle: 'Theme',
                logout: 'Logout',

                server_info: 'Server Information',
                os: 'OS',
                internal_ip: 'Internal IP',
                net_traffic: 'Total Network Traffic',
                total_up: 'Total Upload',
                total_down: 'Total Download',
                
                process: 'Processes',
                network: 'Network',
                service: 'Service',
                connection: 'Connection',
                kill: 'Kill Process',
                restart_term: 'Restart Terminal',

                cron_title: 'Automation (Cron)',
                cron_select: '-- Select cycle --',
                cron_min: 'Every minute',
                cron_5min: 'Every 5 mins',
                cron_hour: 'Every hour',
                cron_day: 'Every 00:00 daily',
                cron_cmd: 'Execution command...',
                cron_empty: 'No Cronjobs yet',

                tg_title: 'Telegram Alerts',
                tg_cond: 'Alert Conditions',
                tg_cpu: 'CPU threshold',
                tg_ram: 'RAM threshold',
                tg_disk: 'Disk full',
                save_cfg: 'Save Configuration',

                cancel: 'Cancel',
                confirm: 'Confirm',

                msg_login_ok: 'Login successful',
                msg_conn_err: 'Server connection error',
                msg_logout_title: 'Logout',
                msg_logout_desc: 'Are you sure you want to log out?',
                msg_cron_add_err: 'Please fill in both schedule and command!',
                msg_cron_added: 'Cronjob added',
                msg_cron_del_title: 'Delete Cronjob',
                msg_cron_del_desc: 'Are you sure you want to delete this task?',
                msg_deleted: 'Deleted successfully',
                msg_cfg_saved: 'Configuration saved',
                msg_cfg_err: 'Error saving configuration',
                msg_kill_title: 'Kill process',
                msg_kill_desc: 'Force stop process {pid} ({name})?',
                msg_kill_ok: 'Killed PID: {pid}',
                msg_kill_err: 'Error killing process',
                msg_term_title: 'Restart Terminal?',
                msg_term_desc: 'This will disconnect the current session and create a new Terminal. Are you sure?',
                msg_term_ok: 'Terminal session refreshed',
                cpu_load: 'CPU Load',
                ram_usage: 'RAM Usage',
                disk_usage: 'Disk Usage',
                clock_speed: 'Clock Speed',
                cores: 'Cores',
                threads: 'Threads',
                active: 'Active',
                cached: 'Cached',
                buffers: 'Buffers',
                ports: 'Ports',
                port: 'Port',
                bind_ip: 'Bind IP',
                cron_15min: 'Every 15 mins',
                cron_30min: 'Every 30 mins',
                cron_12hour: 'Every 12 hours',
                cron_sunday: 'Every Sunday',
                cron_month: '1st of every month',
                tg_test_empty: "Please enter Bot Token and Chat ID before testing!",
                tg_test_sending: "Sending test message...",
                tg_test_success: "Test message sent successfully!",
                tg_test_failed: "Failed to send message",
                error: "Error",
                network_error: "Network error while connecting to server.",
                test_cfg: "Test"
            }
        },

        t(key) {
            try {
                return this.translations[this.lang][key] || key;
            } catch (e) {
                return key;
            }
        },

        toggleLang() {
            this.lang = this.lang === 'vi' ? 'en' : 'vi';
            localStorage.setItem('lang', this.lang);
        },

        isAuthenticated: false,
        showMobileWarning: true,
        loginForm: { username: '', password: '' },
        token: null,
        isDark: localStorage.getItem('theme') === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches),
        isLoadingData: false,
        isLoadingDashboard: false,

        activeTab: 'dashboard',
        monitorSubTab: 'system',
        tabsList: [{
            id: 'dashboard', key: 'tab_dashboard', icon: 'fa-chart-pie'
        }, {
            id: 'monitor', key: 'tab_monitor', icon: 'fa-server'
        }, {
            id: 'terminal', key: 'tab_terminal', icon: 'fa-terminal'
        }, {
            id: 'settings', key: 'tab_settings', icon: 'fa-sliders'
        }],

        currentData: [],
        portData: [],
        cronList: [],
        newCron: { schedule: '', command: '', comment: 'Created via Web' },
        cronPreset: 'custom',
        tgSettings: {
            tg_token: '', tg_chat_id: '',
            tg_cpu_enabled: true, tg_cpu_threshold: 90,
            tg_ram_enabled: true, tg_ram_threshold: 90,
            tg_disk_enabled: true, tg_disk_threshold: 90
        },
        stats: {
            cpu: 0, cpu_per_core: [], cpu_cores: 0, cpu_threads: 0,
            ram: 0, ram_used: 0, ram_total: 0, swap_percent: 0, swap_used: 0, swap_total: 0,
            disks: [], uptime: '0h 0m', sys_info: null, net_io: null
        },

        searchQuery: '',
        sortColumn: 'cpu_percent',
        sortAscending: false,
        fetchInterval: null,
        toasts: [],
        confirmDialog: { show: false, title: '', message: '', callback: null },
        term: null, termWs: null, termInited: false, fitAddon: null,

        init() {
            this.applyTheme();
            this.$watch('isDark', val => {
                localStorage.setItem('theme', val ? 'dark' : 'light');
                this.applyTheme();
            });

            this.token = localStorage.getItem('vps_token');
            if (this.token) {
                this.isAuthenticated = true;
                this.startPolling();
            }
        },

        applyTheme() {
            if (this.isDark) document.documentElement.classList.add('dark');
            else document.documentElement.classList.remove('dark');
        },
        toggleTheme() {
            this.isDark = !this.isDark;
        },

        showToast(message, type = 'info') {
            const id = Date.now();
            this.toasts.push({ id, message, type, show: false });
            setTimeout(() => {
                const t = this.toasts.find(x => x.id === id);
                if (t) t.show = true;
            }, 50);
            setTimeout(() => {
                const t = this.toasts.find(x => x.id === id);
                if (t) t.show = false;
                setTimeout(() => {
                    this.toasts = this.toasts.filter(x => x.id !== id);
                }, 300);
            }, 3000);
        },
        showConfirm(title, message, callback) {
            this.confirmDialog = { show: true, title, message, callback };
        },
        executeConfirm() {
            if (typeof this.confirmDialog.callback === 'function') this.confirmDialog.callback();
            this.confirmDialog.show = false;
        },

        applyCronPreset() {
            if (this.cronPreset !== 'custom') this.newCron.schedule = this.cronPreset;
        },

        async doLogin() {
            try {
                const res = await fetch('/api/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(this.loginForm)
                });
                const data = await res.json();
                if (res.ok) {
                    this.token = data.token;
                    localStorage.setItem('vps_token', this.token);
                    this.isAuthenticated = true;
                    this.startPolling();
                    this.showToast(this.t('msg_login_ok'), 'success');
                } else {
                    this.showToast(data.detail, 'error');
                }
            } catch (err) {
                this.showToast(this.t('msg_conn_err'), 'error');
            }
        },
        confirmLogout() {
            this.showConfirm(this.t('msg_logout_title'), this.t('msg_logout_desc'), () => this.logout());
        },
        logout() {
            this.isAuthenticated = false;
            this.token = null;
            localStorage.removeItem('vps_token');
            clearInterval(this.fetchInterval);
            if (this.termWs) this.termWs.close();
        },

        async apiFetch(endpoint, options = {}) {
            options.headers = {...options.headers, 'Authorization': `Bearer ${this.token}`};
            const res = await fetch(endpoint, options);
            if (res.status === 401) {
                this.logout();
                throw new Error("Phiên đăng nhập hết hạn");
            }
            return res;
        },

        switchTab(tab) {
            this.activeTab = tab;
            if (tab === 'dashboard') this.loadStats(true);
            if (tab === 'monitor') this.setMonitorTab(this.monitorSubTab);
            if (tab === 'settings') {
                this.loadCrons();
                this.loadSettings();
            }
            if (tab === 'terminal' && !this.termInited) this.initTerminal();
            window.scrollTo({ top: 0, behavior: 'smooth' });
        },

        setMonitorTab(subTab) {
            this.monitorSubTab = subTab;
            if (subTab === 'system') {
                this.sortColumn = 'cpu_percent';
                this.sortAscending = false;
                this.loadData();
            }
            if (subTab === 'network') {
                this.sortColumn = 'tx';
                this.sortAscending = false;
                this.loadData();
            }
            if (subTab === 'ports') this.loadPorts();
        },

        startPolling() {
            this.switchTab(this.activeTab);
            this.fetchInterval = setInterval(() => {
                if (this.activeTab === 'dashboard') this.loadStats(false);
                if (this.activeTab === 'monitor' && ['system', 'network'].includes(this.monitorSubTab)) this.loadData(false);
            }, 3000);
        },

        async loadStats(showLoading = true) {
            if (showLoading && !this.stats.cpu) this.isLoadingDashboard = true;
            try {
                const res = await this.apiFetch('/api/stats');
                this.stats = await res.json();
            } catch (e) {} finally { this.isLoadingDashboard = false; }
        },

        async loadData(showLoading = true) {
            if (showLoading && this.currentData.length === 0) this.isLoadingData = true;
            try {
                const res = await this.apiFetch(this.monitorSubTab === 'network' ? '/api/network' : '/api/system');
                this.currentData = await res.json();
            } catch (e) {} finally { this.isLoadingData = false; }
        },
        async loadPorts() {
            this.isLoadingData = true;
            try {
                const res = await this.apiFetch('/api/ports');
                this.portData = await res.json();
            } catch (e) {} finally { this.isLoadingData = false; }
        },

        async loadCrons() {
            try {
                const res = await this.apiFetch('/api/cron');
                this.cronList = await res.json();
            } catch (e) {}
        },
        async saveCron() {
            if (!this.newCron.schedule || !this.newCron.command) return this.showToast(this.t('msg_cron_add_err'), 'error');
            try {
                await this.apiFetch('/api/cron', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(this.newCron)
                });
                this.newCron.schedule = '';
                this.newCron.command = '';
                this.cronPreset = 'custom';
                this.loadCrons();
                this.showToast(this.t('msg_cron_added'), 'success');
            } catch (e) {
                this.showToast('Lỗi khi thêm', 'error');
            }
        },
        deleteCron(id) {
            this.showConfirm(this.t('msg_cron_del_title'), this.t('msg_cron_del_desc'), async() => {
                try {
                    await this.apiFetch(`/api/cron/${id}`, { method: 'DELETE' });
                    this.loadCrons();
                    this.showToast(this.t('msg_deleted'), 'success');
                } catch (e) {}
            });
        },

        async loadSettings() {
            try {
                const res = await this.apiFetch('/api/settings');
                this.tgSettings = await res.json();
            } catch (e) {}
        },
        async saveSettings() {
            try {
                await this.apiFetch('/api/settings', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(this.tgSettings)
                });
                this.showToast(this.t('msg_cfg_saved'), 'success');
            } catch (e) {
                this.showToast(this.t('msg_cfg_err'), 'error');
            }
        },
        
        async testTelegram() {
            if (!this.tgSettings.tg_token || !this.tgSettings.tg_chat_id) {
                this.showToast(this.t('tg_test_empty'), 'error');
                return;
            }
            
            this.showToast(this.t('tg_test_sending'), 'info');
        
            try {
                const res = await fetch('/api/settings/test', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${this.token}`
                    },
                    body: JSON.stringify({
                        tg_token: this.tgSettings.tg_token,
                        tg_chat_id: this.tgSettings.tg_chat_id
                    })
                });
                
                const data = await res.json();
                
                if (res.ok) {
                    this.showToast(this.t('tg_test_success'), 'success');
                } else {
                    this.showToast(`${this.t('error')}: ${data.detail || this.t('tg_test_failed')}`, 'error');
                }
            } catch (err) {
                this.showToast(this.t('network_error'), 'error');
            }
        },

        killProcess(pid, name) {
            let msg = this.t('msg_kill_desc').replace('{pid}', pid).replace('{name}', name);
            this.showConfirm(this.t('msg_kill_title'), msg, async() => {
                try {
                    await this.apiFetch(`/api/kill/${pid}`, { method: 'POST' });
                    this.loadData();
                    this.showToast(this.t('msg_kill_ok').replace('{pid}', pid), 'success');
                } catch (e) {
                    this.showToast(this.t('msg_kill_err'), 'error');
                }
            });
        },

        initTerminal() {
            setTimeout(() => {
                const container = document.getElementById('terminal-container');
                if (!container || this.termInited) return;

                container.innerHTML = '';

                const term = new Terminal({
                    theme: { background: 'transparent', foreground: '#e2e8f0', cursor: '#3b82f6' },
                    cursorBlink: true, fontSize: 14, fontFamily: 'Menlo, Monaco, "Courier New", monospace'
                });
                const fitAddon = new FitAddon.FitAddon();

                term.loadAddon(fitAddon);
                term.open(container);
                fitAddon.fit();

                const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
                try {
                    this.termWs = new WebSocket(`${protocol}//${window.location.host}/ws/terminal?token=${this.token}`);
                    this.termWs.onopen = () => this.termWs.send(`__RESIZE__:${term.cols},${term.rows}`);
                    this.termWs.onmessage = (ev) => term.write(ev.data);
                    term.onData(data => { if (this.termWs.readyState === 1) this.termWs.send(data); });
                    term.onResize((size) => { if (this.termWs && this.termWs.readyState === 1) this.termWs.send(`__RESIZE__:${size.cols},${size.rows}`); });
                } catch (e) {
                    term.write("Lỗi kết nối Terminal.\r\n");
                }

                this.term = term;
                this.fitAddon = fitAddon;

                if (!window.termResizeHandler) {
                    window.termResizeHandler = () => {
                        if (this.fitAddon) {
                            try { Alpine.raw(this.fitAddon).fit(); } catch (e) {}
                        }
                    };
                    window.addEventListener('resize', window.termResizeHandler);
                }

                this.termInited = true;
            }, 200);
        },

        confirmRefreshTerminal() {
            this.showConfirm(this.t('msg_term_title'), this.t('msg_term_desc'), () => this.refreshTerminal());
        },

        refreshTerminal() {
            if (this.termWs) {
                this.termWs.close();
                this.termWs = null;
            }
            if (this.term) {
                try { Alpine.raw(this.term).dispose(); } catch (e) { console.warn(e); }
                this.term = null;
            }
            if (this.fitAddon) this.fitAddon = null;
            this.termInited = false;

            setTimeout(() => {
                this.initTerminal();
                this.showToast(this.t('msg_term_ok'), 'success');
            }, 300);
        },

        sortBy(col) {
            if (this.sortColumn === col) {
                this.sortAscending = !this.sortAscending;
            } else {
                this.sortColumn = col;
                this.sortAscending = false;
            }
        },

        get filteredAndSortedData() {
            return this.currentData.filter(item => {
                const search = this.searchQuery.toLowerCase();
                return String(item.pid).includes(search) || String(item.name).toLowerCase().includes(search);
            }).sort((a, b) => {
                let valA = a[this.sortColumn];
                let valB = b[this.sortColumn];
                if (['cpu_percent', 'memory_percent', 'tx', 'rx', 'pid'].includes(this.sortColumn)) {
                    valA = parseFloat(valA) || 0;
                    valB = parseFloat(valB) || 0;
                } else {
                    valA = String(valA || '').toLowerCase();
                    valB = String(valB || '').toLowerCase();
                }
                if (valA < valB) return this.sortAscending ? -1 : 1;
                if (valA > valB) return this.sortAscending ? 1 : -1;
                return 0;
            });
        }
    }));
});