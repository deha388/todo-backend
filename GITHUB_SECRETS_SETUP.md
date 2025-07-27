# 🔐 GitHub Secrets Setup Guide

**Required for Automated Deployment Pipeline**

## 📋 Steps to Configure Secrets

### **1️⃣ Navigate to Repository Settings**
```
https://github.com/deha388/todo-backend
→ Settings (tab)
→ Secrets and variables (left sidebar)
→ Actions
→ New repository secret
```

### **2️⃣ Add Required Secrets**

#### **Secret 1: DO_HOST**
```
Name: DO_HOST
Value: 142.93.173.157
Description: DigitalOcean VPS IP address
```

#### **Secret 2: DO_USERNAME**
```
Name: DO_USERNAME  
Value: root
Description: VPS SSH username
```

#### **Secret 3: DO_PASSWORD**
```
Name: DO_PASSWORD
Value: [Your VPS root password]
Description: VPS SSH password
```

## ✅ Verification

After adding secrets, you should see:

```
Repository secrets (3)
├── DO_HOST ●●●●●●●●●●●●●●●●●
├── DO_USERNAME ●●●●
└── DO_PASSWORD ●●●●●●●●●●●●●●●●
```

## 🚀 Test Deployment

### **Trigger First Deployment**
```bash
# Any push to main will trigger deployment
git push origin main

# Or manually trigger via GitHub Actions:
# Actions tab → Go Backend Deploy Pipeline → Run workflow
```

### **Monitor Deployment**
```
GitHub Repository
→ Actions tab  
→ Latest workflow run
→ Watch deployment logs
```

### **Expected Deployment Steps**
```
✅ Checkout code
✅ Setup Go 1.21
✅ Download dependencies  
✅ Run all tests (unit + integration + contract)
✅ Build Go application
✅ SSH to VPS
✅ Build Docker image
✅ Import to K3s
✅ Apply Kubernetes deployment
✅ Wait for rollout
✅ Health check verification
```

## 🔍 Troubleshooting

### **Common Issues**

#### **SSH Connection Failed**
- ✅ Check `DO_HOST` is correct IP
- ✅ Check `DO_USERNAME` is `root`  
- ✅ Check `DO_PASSWORD` is correct
- ✅ Verify VPS is running

#### **Docker Build Failed**
- ✅ Check Go code compiles locally
- ✅ Check all tests pass locally
- ✅ Verify Dockerfile syntax

#### **K3s Deployment Failed**
- ✅ Check VPS has K3s installed
- ✅ Check sufficient resources (memory/CPU)
- ✅ Verify `k8s-deployment.yaml` syntax

### **Deployment Logs**
```bash
# On VPS, check logs:
k3s kubectl get pods
k3s kubectl logs -l app=todo-backend
k3s kubectl describe pod <pod-name>
```

## 📞 Support

If deployment fails:
1. Check GitHub Actions logs first
2. Verify all 3 secrets are correctly set
3. Test VPS SSH connection manually
4. Create GitHub issue with error logs

---

**Ready for automated deployment!** 🚀 