import unittest
import subprocess
import sys
import os

SCRIPT = os.path.join(os.path.dirname(__file__), "conventional_commit.py")

class TestConventionalCommit(unittest.TestCase):
    def run_script(self, summary):
        result = subprocess.run([sys.executable, SCRIPT, summary], capture_output=True, text=True)
        self.assertEqual(result.returncode, 0, f"Non-zero exit: {result.stderr}")
        return result.stdout.strip()

    def test_fix_type(self):
        msg = self.run_script("fix bug in user login logic")
        self.assertTrue(msg.startswith("fix:"))

    def test_feat_type(self):
        msg = self.run_script("add new payment method")
        self.assertTrue(msg.startswith("feat:"))

    def test_docs_type(self):
        msg = self.run_script("docs: update API docs")
        self.assertTrue(msg.startswith("docs:"))

    def test_scope_paren(self):
        msg = self.run_script("fix(auth): correct token refresh")
        self.assertEqual(msg, "fix(auth): correct token refresh")

    def test_scope_bracket(self):
        msg = self.run_script("feat [user-profile]: add avatar upload")
        self.assertEqual(msg, "feat(user-profile): add avatar upload")

    def test_header_truncation(self):
        long_summary = "feat: " + "a" * 100
        msg = self.run_script(long_summary)
        self.assertTrue(len(msg) <= 72)
        self.assertTrue(msg.endswith("..."))

    def test_default_type(self):
        msg = self.run_script("update dependencies")
        self.assertTrue(msg.startswith("chore:"))

if __name__ == "__main__":
    unittest.main()
